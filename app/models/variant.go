package models

import (
	"gorm.io/datatypes"
)

type VariantType struct {
	BaseModel

	Name string `gorm:"unique;not null" json:"name"`
}

type Variant struct {
	BaseModel

	Name          string         `gorm:"unique;not null" json:"name"`
	VariantTypeID uint           `json:"variant_type_id"`
	VariantType   VariantType    `gorm:"foreignKey:VariantTypeID" json:"variant_type"`
	Values        datatypes.JSON `gorm:"type:json" json:"values"`
}
