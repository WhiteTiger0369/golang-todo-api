package auth

import (
	"ex1/todo-api/common"
	"ex1/todo-api/pkg"
	"ex1/todo-api/user"
	"net/http"
)

type ServiceAuth interface {
	Login(user user.User) (user.User, common.DatabaseError)
}
type authService struct {
	userRepository user.RepositoryUser
}

func ProvideAuthService(u user.RepositoryUser) *authService {
	return &authService{userRepository: u}
}

func (a *authService) Login(user user.User) (user.User, common.DatabaseError) {

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
