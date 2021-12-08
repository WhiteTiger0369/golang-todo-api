package services

import (
	"ex1/todo-api/common"
	"ex1/todo-api/entities"
	"ex1/todo-api/pkg"
	"ex1/todo-api/repositories"
	"net/http"
)

type AuthService interface {
	Login(user entities.User) (entities.User, common.DatabaseError)
}

type authService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(u repositories.UserRepository) *authService {
	return &authService{userRepository: u}
}

func (a *authService) Login(user entities.User) (entities.User, common.DatabaseError) {

	errorCode := common.DatabaseError{}
	checkUser, err := a.userRepository.FindByUserName(user.Username)
	if err.Type == "error_01" {
		errorCode = common.DatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return user, errorCode
	}

	comparePassword := pkg.ComparePassword(checkUser.Password, user.Password)

	if comparePassword != nil {
		errorCode = common.DatabaseError{
			Code: http.StatusBadRequest,
			Type: "error_02",
		}
		return user, errorCode
	}

	return checkUser, errorCode
}
