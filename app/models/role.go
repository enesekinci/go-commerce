package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name  string `gorm:"not null" json:"name"`
	Slug  string `gorm:"not null" json:"slug"`
	Users []User `gorm:"foreignKey:RoleID"`
}
