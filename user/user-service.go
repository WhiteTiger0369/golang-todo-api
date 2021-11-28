package user

type UserService struct {
	UserRepository UserRepository
}

func ProvideUserService(t UserRepository) UserService {
	return UserService{UserRepository: t}
}

func (u *UserService) FindAll() []User {
	return u.UserRepository.FindAll()
}

func (u *UserService) FindByID(id uint) User {
	return u.UserRepository.FindByID(id)
}

func (u *UserService) Save(user User) User {
	u.UserRepository.Save(user)

	return user
}

func (u *UserService) Delete(user User) {
	u.UserRepository.Delete(user)
}
