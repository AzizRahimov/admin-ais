package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Code string `json:"code"`
	Section string `json:"section"`
	Description string `json:"description"`
	IsRemoved bool `json:"is_removed"`

}
