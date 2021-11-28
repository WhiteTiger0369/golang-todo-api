package user

import "ex1/todo-api/common"

type UserService struct {
	UserRepository UserRepository
}

func ProvideUserService(t UserRepository) UserService {
	return UserService{UserRepository: t}
}

func (u *UserService) FindAll() ([]UserDTO, common.DatabaseError) {
	res, err := u.UserRepository.FindAll()
	return ToUserDTOs(res), err
}

func (u *UserService) FindByID(id uint) (UserDTO, common.DatabaseError) {
	res, err := u.UserRepository.FindByID(id)
	return ToUserDTO(res), err
}

func (u *UserService) Save(userDTO UserDTO) (UserDTO, common.DatabaseError) {
	res, err := u.UserRepository.Save(ToUser(userDTO))
	return ToUserDTO(res), err
}

func (u *UserService) Delete(id uint) {
	u.UserRepository.Delete(id)
}
