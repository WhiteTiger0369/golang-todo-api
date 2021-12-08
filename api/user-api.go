package api

import (
	"ex1/todo-api/dtos"
	"ex1/todo-api/helpers"
	"ex1/todo-api/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type userAPI struct {
	userService services.UserService
}

func ProvideUserAPI(u services.UserService) *userAPI {
	return &userAPI{userService: u}
}

func (u *userAPI) FindAll(c *gin.Context) {
	res, err := u.userService.FindAll()

	switch err.Type {
	case "error_01":
		helpers.APIResponse(c, "Users data is not exists", err.Code, http.MethodGet, nil)
	default:
		helpers.APIResponse(c, "Results Users data successfully", http.StatusOK, http.MethodGet, res)
	}
}

func (u *userAPI) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := u.userService.FindByID(uint(id))

	switch err.Type {
	case "error_01":
		helpers.APIResponse(c, "Users data is not exists", err.Code, http.MethodGet, nil)
	default:
		helpers.APIResponse(c, "Results Users data successfully", http.StatusOK, http.MethodGet, res)
	}
}

func (u *userAPI) Create(c *gin.Context) {
	var userDTO dtos.UserDTO
	errIn := c.BindJSON(&userDTO)
	if errIn != nil {
		c.Status(http.StatusBadRequest)
		log.Fatalln(errIn)
		return
	}

	res, err := u.userService.Save(userDTO)

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

func (u *userAPI) Update(c *gin.Context) {
	var userDTO dtos.UserDTO
	err := c.BindJSON(&userDTO)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Fatalln(err)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	user, _ := u.userService.FindByID(uint(id))
	if user == (dtos.UserDTO{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	user.FullName = userDTO.FullName
	user.Username = userDTO.Username
	user.Password = userDTO.Password
	u.userService.Save(user)

	c.Status(http.StatusOK)
}

func (u *userAPI) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	u.userService.Delete(uint(id))

	c.Status(http.StatusOK)
}
