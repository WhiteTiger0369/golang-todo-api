package user

import "github.com/jinzhu/gorm"

type UserRepository struct {
	DB *gorm.DB
}

func ProvideUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}

func (u *UserRepository) FindAll() []User {
	var user []User
	u.DB.Find(&user)

	return user
}

func (u *UserRepository) FindByID(id uint) User {
	var user User
	u.DB.First(&user, id)

	return user
}

func (u *UserRepository) Save(user User) User {
	u.DB.Save(&user)

	return user
}

func (u *UserRepository) Delete(user User) {
	u.DB.Delete(&user)
}
