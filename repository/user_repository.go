package repository

import (
	"admin-ais/model"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{DB: db}
}

// RoleRepository : represent the user's repository contract
type UserRepository interface {
	AddUser(model.User) (model.User, error)
	GetAllUsers() ([]model.User, error)
	ChangeUser(model.User, string) error
	DeleteUser(id string) error
	Migrate() error
}

func (u userRepository) AddUser(user model.User) (model.User, error) {
	return user, u.DB.Create(&user).Error

}

func (u userRepository) GetAllUsers() ([]model.User,  error) {
	var users []model.User


	return users, u.DB.Where("is_removed = ?", "false").Find(&users).Error
}

func (u userRepository) ChangeUser(inputUser model.User, id string) error {
	// проверка сущ ли пользователь
	var user model.User
	err := u.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}

	return u.DB.Model(&user).Where("id = ?", id).Updates(inputUser).Error
}

func (u userRepository) DeleteUser(id string) error {
	var user model.User
	err := u.DB.First(&user, id).Error
	if err != nil {
		return err
	}

//	return r.DB.Model(right).Where("id = ?", id).Update("is_removed", "true").Error

	return u.DB.Model(&user).Update("is_removed", true).Error

}

func (u userRepository) Migrate() error {
	err := u.DB.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}

	return nil
}



//func (r rightRepository) Migrate() error {
//	err := r.DB.AutoMigrate(&model.User{})
//	if err != nil {
//		return err
//	}
//
//	return err
//}
//
//func (r rightRepository) AddRight(right model.Right) (model.Right, error) {
//
//	return right, r.DB.Create(&right).Error
//
//}
//
//func (r rightRepository) GetAllRights() (right []model.Right, err error) {
//
//	return right, r.DB.Where("is_removed", false).Find(&right).Error
//
//}
//
//
//
//func (r rightRepository) ChangeRight(inputRight model.Right, id string) error {
//	var right model.Right
//
//	return r.DB.Model(&right).Where("id = ?", id).Updates(inputRight).Error
//
//}
//
//func (r rightRepository) DeleteRight(right model.Right, id string) error {
//	err := r.DB.First(&right, id).Error
//	if err != nil {
//
//		return errors.New("Role not found")
//	}
//	return r.DB.Model(right).Where("id = ?", id).Update("is_removed", "true").Error
//
//}
//
