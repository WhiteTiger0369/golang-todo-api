package api

import (
	"ex1/todo-api/entities"
	"ex1/todo-api/helpers"
	"ex1/todo-api/pkg"
	"ex1/todo-api/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

type authAPI struct {
	authService services.AuthService
}

func NewAuthAPI(a services.AuthService) *authAPI {
	return &authAPI{authService: a}
}

func (u *authAPI) Login(c *gin.Context) {
	var userReq entities.User
	errIn := c.BindJSON(&userReq)
	if errIn != nil {
		c.Status(http.StatusBadRequest)
		log.Fatalln(errIn)
		return
	}

	checkUser, err := u.authService.Login(userReq)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(c, "User account is not registered", err.Code, http.MethodPost, nil)
		return
	case "error_02":
		helpers.APIResponse(c, "Username or password is wrong", err.Code, http.MethodPost, nil)
		return
	default:
		accessTokenData := map[string]interface{}{"username": checkUser.Username, "password": checkUser.Password, "id": checkUser.ID}
		accessToken, errToken := pkg.Sign(accessTokenData, "JWT_SECRET", 24*60*1)

		if errToken != nil {
			defer logrus.Error(errToken.Error())
			helpers.APIResponse(c, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
			return
		}

		helpers.APIResponse(c, "Login successfully", http.StatusOK, http.MethodPost, map[string]string{"accessToken": accessToken})
	}

}
