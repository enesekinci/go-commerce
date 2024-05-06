package models

import "gorm.io/gorm"

type Tax struct {
	gorm.Model

	Name string  `gorm:"not null;unique" json:"name"`
	Rate float64 `gorm:"default:0" json:"rate"`
}

type TaxClass struct {
	gorm.Model
}

// sistemin genel kullanacagı modeller burada yer alacak.

type Setting struct {
	gorm.Model
	Key   string `gorm:"not null;unique" json:"name"`
	Value string `gorm:"type:text;nullable" json:"value"`
}

type Media struct {
	gorm.Model
	Name string `gorm:"not null" json:"name"`
	Path string `gorm:"not null" json:"path"`
	Type string `gorm:"not null" json:"type"`
}

type Language struct {
	gorm.Model
	Name      string `gorm:"not null;unique" json:"name"`
	Code      string `gorm:"not null;unique" json:"code"`
	Status    uint   `gorm:"default:1" json:"status"`
	IsCurrent uint   `gorm:"default:0" json:"is_current"`
}

type Currency struct {
	gorm.Model
	Name   string  `gorm:"not null;unique" json:"name"`
	Code   string  `gorm:"not null;unique" json:"code"`
	Symbol string  `gorm:"not null" json:"symbol"`
	Rate   float64 `gorm:"default:1" json:"rate"`
	Status uint    `gorm:"default:1" json:"status"`
}

type Country struct {
	gorm.Model
	Name   string `gorm:"not null;unique" json:"name"`
	Code   string `gorm:"not null;unique" json:"code"`
	Status uint   `gorm:"default:1" json:"status"`
}

type Zone struct {
	gorm.Model
	Name      string `gorm:"not null;unique" json:"name"`
	Code      string `gorm:"not null;unique" json:"code"`
	Country   Country
	CountryID uint `json:"country_id"`
	Status    uint `gorm:"default:1" json:"status"`
}

type PaymentMethod struct {
	gorm.Model
	Name   string `gorm:"not null;unique" json:"name"`
	Status uint   `gorm:"default:1" json:"status"`
}

type ShippingMethod struct {
	gorm.Model
	Name   string `gorm:"not null;unique" json:"name"`
	Status uint   `gorm:"default:1" json:"status"`
}
