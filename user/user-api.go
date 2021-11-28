package user

import (
	"ex1/todo-api/helpers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type UserAPI struct {
	UserService UserService
}

func ProvideUserAPI(u UserService) UserAPI {
	return UserAPI{UserService: u}
}

func (u *UserAPI) FindAll(c *gin.Context) {
	res, err := u.UserService.FindAll()

	switch err.Type {
	case "error_01":
		helpers.APIResponse(c, "Users data is not exists", err.Code, http.MethodGet, nil)
	default:
		helpers.APIResponse(c, "Results Users data successfully", http.StatusOK, http.MethodGet, res)
	}
}

func (u *UserAPI) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := u.UserService.FindByID(uint(id))

	switch err.Type {
	case "error_01":
		helpers.APIResponse(c, "Users data is not exists", err.Code, http.MethodGet, nil)
	default:
		helpers.APIResponse(c, "Results Users data successfully", http.StatusOK, http.MethodGet, res)
	}
}

func (u *UserAPI) Create(c *gin.Context) {
	var userDTO UserDTO
	errIn := c.BindJSON(&userDTO)
	if errIn != nil {
		c.Status(http.StatusBadRequest)
		log.Fatalln(errIn)
		return
	}

	res, err := u.UserService.Save(userDTO)

	switch err.Type {
	case "error_01":
		helpers.APIResponse(c, "User already exist", err.Code, http.MethodPost, nil)
		return
	case "error_02":
		helpers.APIResponse(c, "Create new user account failed", err.Code, http.MethodPost, nil)
		return
	default:
		helpers.APIResponse(c, "Create new user account successfully", http.StatusCreated, http.MethodPost, res)
	}
}

func (u *UserAPI) Update(c *gin.Context) {
	var userDTO UserDTO
	err := c.BindJSON(&userDTO)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Fatalln(err)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	user, _ := u.UserService.FindByID(uint(id))
	if user == (UserDTO{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	user.FullName = userDTO.FullName
	user.Username = userDTO.Username
	user.Password = userDTO.Password
	u.UserService.Save(user)

	c.Status(http.StatusOK)
}

func (u *UserAPI) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	u.UserService.Delete(uint(id))

	c.Status(http.StatusOK)
}
