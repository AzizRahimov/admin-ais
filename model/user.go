package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name"`
	IsRemoved bool `json:"is_removed"`
}

type UserRoles struct {
	UserID  User
	RoleID Role

}
type UserRights struct {
	UserID  User
	RightID Right

}