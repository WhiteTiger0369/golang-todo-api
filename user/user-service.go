package user

import "ex1/todo-api/common"

type ServiceUser interface {
	FindAll() ([]UserDTO, common.DatabaseError)
	FindByID(id uint) (UserDTO, common.DatabaseError)
	Save(userDTO UserDTO) (UserDTO, common.DatabaseError)
	Delete(id uint)
}
type userService struct {
	userRepository RepositoryUser
}

func ProvideUserService(u RepositoryUser) *userService {
	return &userService{userRepository: u}
}

func (u *userService) FindAll() ([]UserDTO, common.DatabaseError) {
	res, err := u.userRepository.FindAll()
	return ToUserDTOs(res), err
}

func (u *userService) FindByID(id uint) (UserDTO, common.DatabaseError) {
	res, err := u.userRepository.FindByID(id)
	return ToUserDTO(res), err
}

func (u *userService) Save(userDTO UserDTO) (UserDTO, common.DatabaseError) {
	res, err := u.userRepository.Save(ToUser(userDTO))
	return ToUserDTO(res), err
}

func (u *userService) Delete(id uint) {
	u.userRepository.Delete(id)
}
