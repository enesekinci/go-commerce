package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Attribute struct {
	BaseModel
	Name       string         `gorm:"not null" json:"name"`
	Filterable uint           `gorm:"default:0" json:"filterable"`
	Values     datatypes.JSON `gorm:"type:json" json:"values"`
	Categories []Category     `gorm:"many2many:attribute_categories;" json:"categories"`
}

type AttributeCategory struct {
	BaseModel
	AttributeID uint `json:"attribute_id"`
	CategoryID  uint `json:"category_id"`

	Attribute Attribute `gorm:"foreignKey:AttributeID;" json:"attribute"`
	Category  Category  `gorm:"foreignKey:CategoryID;" json:"category"`
}

func NewAttribute(name string, filterable uint, values string) Attribute {
	return Attribute{
		Name:       name,
		Filterable: filterable,
		Values:     []byte(values),
		Categories: []Category{},
	}
}

func (a *Attribute) AfterDelete(tx *gorm.DB) (err error) {
	tx.Where("attribute_id = ?", a.ID).Delete(&AttributeCategory{})
	return
}
