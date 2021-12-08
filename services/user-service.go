package services

import (
	"ex1/todo-api/common"
	"ex1/todo-api/dtos"
	"ex1/todo-api/mappers"
	"ex1/todo-api/repositories"
)

type UserService interface {
	FindAll() ([]dtos.UserDTO, common.DatabaseError)
	FindByID(id uint) (dtos.UserDTO, common.DatabaseError)
	Save(userDTO dtos.UserDTO) (dtos.UserDTO, common.DatabaseError)
	Delete(id uint)
}
type userService struct {
	userRepository repositories.RepositoryUser
}

func ProvideUserService(u repositories.RepositoryUser) *userService {
	return &userService{userRepository: u}
}

func (u *userService) FindAll() ([]dtos.UserDTO, common.DatabaseError) {
	res, err := u.userRepository.FindAll()
	return mappers.ToUserDTOs(res), err
}

func (u *userService) FindByID(id uint) (dtos.UserDTO, common.DatabaseError) {
	res, err := u.userRepository.FindByID(id)
	return mappers.ToUserDTO(res), err
}

func (u *userService) Save(userDTO dtos.UserDTO) (dtos.UserDTO, common.DatabaseError) {
	res, err := u.userRepository.Save(mappers.ToUser(userDTO))
	return mappers.ToUserDTO(res), err
}

func (u *userService) Delete(id uint) {
	u.userRepository.Delete(id)
}
