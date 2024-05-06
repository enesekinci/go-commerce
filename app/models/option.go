package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Name,Type (Field,Textarea,Dropdown,Checkbox,Radio,Multiselect,Date,DateTime,Time),Required,Values

type OptionType struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null;unique" json:"name"`
}
type Option struct {
	gorm.Model
	Name   string         `gorm:"type:varchar(255);not null;unique" json:"name"`
	TypeId uint           `gorm:"type:int;not null" json:"type_id"`
	Type   OptionType     `gorm:"foreignKey:TypeId" json:"type"`
	Values datatypes.JSON `gorm:"type:json" json:"values"`
}

func NewOption(name string, typeId uint, values string) Option {
	return Option{
		Name:   name,
		TypeId: typeId,
		Values: []byte(values),
	}
}
