package auth

import (
	"ex1/todo-api/common"
	"ex1/todo-api/pkg"
	"ex1/todo-api/user"
	"net/http"
)

type AuthService struct {
	UserRepository user.UserRepository
}

func ProvideAuthService(t user.UserRepository) AuthService {
	return AuthService{UserRepository: t}
}

func (a *AuthService) Login(user user.User) (user.User, common.DatabaseError) {

	errorCode := common.DatabaseError{}
	checkUser, err := a.UserRepository.FindByUserName(user.Username)
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
