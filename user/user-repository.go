package user

import (
	"ex1/todo-api/common"
	"github.com/jinzhu/gorm"
	"net/http"
)

type UserRepository struct {
	DB *gorm.DB
}

func ProvideUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}

func (u *UserRepository) FindAll() ([]User, common.DatabaseError) {
	var users []User
	errorCode := common.DatabaseError{}

	results := u.DB.Debug().Find(&users)

	if results.RowsAffected < 1 {
		errorCode = common.DatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return users, errorCode
	}

	return users, errorCode
}

func (u *UserRepository) FindByID(id uint) (User, common.DatabaseError) {
	var user User

	errorCode := common.DatabaseError{}

	resultUsers := u.DB.First(&user, id)

	if resultUsers.RowsAffected < 1 {
		errorCode = common.DatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return user, errorCode
	}

	return user, errorCode

}

func (u *UserRepository) Save(user User) (User, common.DatabaseError) {

	errorCode := common.DatabaseError{}

	checkUserExist, _ := u.FindByUserName(user.Username)

	if checkUserExist.ID > 0 {
		errorCode = common.DatabaseError{
			Code: http.StatusForbidden,
			Type: "error_01",
		}
		return user, errorCode
	}

	addUser := u.DB.Debug().Save(&user)
	if addUser.RowsAffected < 1 {
		errorCode = common.DatabaseError{
			Code: http.StatusForbidden,
			Type: "error_02",
		}
		return user, errorCode
	}

	return user, errorCode
}

func (u *UserRepository) Delete(id uint) {
	u.DB.Delete(User{}, "id = ?", id)

}

func (u *UserRepository) FindByUserName(username string) (User, common.DatabaseError) {
	var user User
	errorCode := common.DatabaseError{}
	res := u.DB.First(&user, "username = ?", username)

	if res.RowsAffected < 1 {
		errorCode = common.DatabaseError{
			Code: http.StatusForbidden,
			Type: "error_01",
		}
		return user, errorCode
	}
	return user, errorCode
}
