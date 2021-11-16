package repository

import (
	"admin-ais/model"
	"errors"
	"gorm.io/gorm"
)

type rightRepository struct {
	DB *gorm.DB
}

func NewRightRepository(db *gorm.DB) RightRepository {
	return rightRepository{DB: db}
}

// RoleRepository : represent the user's repository contract
type RightRepository interface {
	AddRight(model.Right) (model.Right, error)
	GetAllRights() (rights []model.Right, err error)
	ChangeRight(model.Right, string) error
	DeleteRight(id string) error
	Migrate() error
}

func (r rightRepository) Migrate() error {
	err := r.DB.AutoMigrate(&model.Right{})
	if err != nil {
		return err
	}

	return err
}

func (r rightRepository) AddRight(right model.Right) (model.Right, error) {

	return right, r.DB.Create(&right).Error

}

func (r rightRepository) GetAllRights() (right []model.Right, err error) {

	return right, r.DB.Where("is_removed", false).Find(&right).Error

}


//func UpdateBook(c *gin.Context) {
func (r rightRepository) ChangeRight(inputRight model.Right, id string) error {
	var right model.Right

	return r.DB.Model(&right).Where("id = ?", id).Updates(inputRight).Error

}

func (r rightRepository) DeleteRight( id string) error {
	var right model.Right

	err := r.DB.First(&right, id).Error
	if err != nil {

		return errors.New("Right not found")
	}
	return r.DB.Model(right).Where("id = ?", id).Update("is_removed", "true").Error

}
