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
	Update(id uint, userDTO dtos.UserDTO) (dtos.UserDTO, common.DatabaseError)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(u repositories.UserRepository) *userService {
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

func (u *userService) Update(id uint, userDTO dtos.UserDTO) (dtos.UserDTO, common.DatabaseError) {
	existsUser, err := u.userRepository.FindByID(id)
	if err.Type == "error_01" {
		return userDTO, err
	}
	existsUser.FullName = userDTO.FullName
	existsUser.Password = userDTO.Password
	res, err := u.userRepository.Save(existsUser)
	return mappers.ToUserDTO(res), err
}
