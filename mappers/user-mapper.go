package mappers

import (
	"ex1/todo-api/dtos"
	"ex1/todo-api/entities"
)

func ToUser(userDTO dtos.UserDTO) entities.User {
	return entities.User{FullName: userDTO.FullName, Username: userDTO.Username, Password: userDTO.Password}
}

func ToUserDTO(user entities.User) dtos.UserDTO {
	return dtos.UserDTO{ID: user.ID, FullName: user.FullName, Username: user.Username, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt}
}

func ToUserDTOs(users []entities.User) []dtos.UserDTO {
	userDTOs := make([]dtos.UserDTO, len(users))

	for i, itm := range users {
		userDTOs[i] = ToUserDTO(itm)
	}

	return userDTOs
}
