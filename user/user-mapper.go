package user

func ToUser(userDTO UserDTO) User {
	return User{FullName: userDTO.FullName, Username: userDTO.Username, Password: userDTO.Password}
}

func ToUserDTO(user User) UserDTO {
	return UserDTO{ID: user.ID, FullName: user.FullName, Username: user.Username, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt}
}

func ToUserDTOs(users []User) []UserDTO {
	userDTOs := make([]UserDTO, len(users))

	for i, itm := range users {
		userDTOs[i] = ToUserDTO(itm)
	}

	return userDTOs
}
