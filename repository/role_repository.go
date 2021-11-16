package repository

import (
	"admin-ais/model"
	"errors"
	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return roleRepository{db: db}
}

// RoleRepository : represent the user's repository contract
type RoleRepository interface {
	AddRole(model.Role) (model.Role, error)
	GetAllRoles() (roles []model.Role, err error)
	ChangeRole(model.Role, string) error
	DeleteRole(role model.Role, id string) error
	Migrate() error
}

func (r roleRepository) Migrate() error {
	err := r.db.AutoMigrate(&model.Role{})
	if err != nil {
		return err
	}

	return err
}

func (r roleRepository) AddRole(role model.Role) (model.Role, error) {
	return role, r.db.Create(&role).Error

}

func (r roleRepository) GetAllRoles() ([]model.Role, error) {
	var roles []model.Role

	err := r.db.Where("is_removed = ?", false).Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return roles, nil

}

func (r roleRepository) ChangeRole(inputRole model.Role, id string) error {
	// проверка сущ ли роль
	var role model.Role

	err := r.db.Where("id = ?", id).First(&role).Error
	if err != nil {
		return err
	}

	return r.db.Model(&role).Updates(inputRole).Error
}

func (r roleRepository) DeleteRole(role model.Role, id string) error {
	err := r.db.First(&role, id).Error
	if err != nil {
		return errors.New("Role not found")
	}


	return r.db.Model(role).Where("id = ?", id).Update("is_removed", "true").Error

}
