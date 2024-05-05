package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name string `validate:"required;unique" json:"name"`
}

func NewTag(name string) *Tag {
	return &Tag{
		Name: name,
	}
}
