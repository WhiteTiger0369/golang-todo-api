package auth

import (
	"ex1/todo-api/helpers"
	"ex1/todo-api/pkg"
	"ex1/todo-api/user"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

type AuthAPI struct {
	AuthService AuthService
}

func ProvideAuthAPI(a AuthService) AuthAPI {
	return AuthAPI{AuthService: a}
}

func (u *AuthAPI) Login(c *gin.Context) {
	var userReq user.User
	errIn := c.BindJSON(&userReq)
	if errIn != nil {
		c.Status(http.StatusBadRequest)
		log.Fatalln(errIn)
		return
	}

	checkUser, err := u.AuthService.Login(userReq)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(c, "User account is not registered", err.Code, http.MethodPost, nil)
		return
	case "error_02":
		helpers.APIResponse(c, "Username or password is wrong", err.Code, http.MethodPost, nil)
		return
	default:
		accessTokenData := map[string]interface{}{"username": checkUser.Username, "password": checkUser.Password}
		accessToken, errToken := pkg.Sign(accessTokenData, "JWT_SECRET", 24*60*1)

		if errToken != nil {
			defer logrus.Error(errToken.Error())
			helpers.APIResponse(c, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
			return
		}

		helpers.APIResponse(c, "Login successfully", http.StatusOK, http.MethodPost, map[string]string{"accessToken": accessToken})
	}

}
