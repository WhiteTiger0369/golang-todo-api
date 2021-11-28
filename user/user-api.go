package user

import (
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
	Users := u.UserService.FindAll()

	c.JSON(http.StatusOK, gin.H{"Users": ToUserDTOs(Users)})
}

func (u *UserAPI) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	User := u.UserService.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"User": ToUserDTO(User)})
}

func (u *UserAPI) Create(c *gin.Context) {
	var userDTO UserDTO
	err := c.BindJSON(&userDTO)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Fatalln(err)
		return
	}

	createdUser := u.UserService.Save(ToUser(userDTO))

	c.JSON(http.StatusOK, gin.H{"User": ToUserDTO(createdUser)})
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
	user := u.UserService.FindByID(uint(id))
	if user == (User{}) {
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
	user := u.UserService.FindByID(uint(id))
	if user == (User{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	u.UserService.Delete(user)

	c.Status(http.StatusOK)
}
