package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Username string
	Password string
	Name     string
	IsAdmin  bool
}
