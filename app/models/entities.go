package models

type Tax struct {
	BaseModel

	Name string  `gorm:"not null;unique" json:"name"`
	Rate float64 `gorm:"default:0" json:"rate"`
}

type TaxClass struct {
	BaseModel
}

// sistemin genel kullanacagı modeller burada yer alacak.

type Setting struct {
	BaseModel
	Key   string `gorm:"not null;unique" json:"name"`
	Value string `gorm:"type:text;nullable" json:"value"`
}

type Media struct {
	BaseModel
	Name string `gorm:"not null" json:"name"`
	Path string `gorm:"not null" json:"path"`
	Type string `gorm:"not null" json:"type"`
}

type Language struct {
	BaseModel
	Name      string `gorm:"not null;unique" json:"name"`
	Code      string `gorm:"not null;unique" json:"code"`
	Status    uint   `gorm:"default:1" json:"status"`
	IsCurrent uint   `gorm:"default:0" json:"is_current"`
}

type Currency struct {
	BaseModel
	Name   string  `gorm:"not null;unique" json:"name"`
	Code   string  `gorm:"not null;unique" json:"code"`
	Symbol string  `gorm:"not null" json:"symbol"`
	Rate   float64 `gorm:"default:1" json:"rate"`
	Status uint    `gorm:"default:1" json:"status"`
}

type Country struct {
	BaseModel
	Name   string `gorm:"not null;unique" json:"name"`
	Code   string `gorm:"not null;unique" json:"code"`
	Status uint   `gorm:"default:1" json:"status"`
}

type Zone struct {
	BaseModel
	Name      string `gorm:"not null;unique" json:"name"`
	Code      string `gorm:"not null;unique" json:"code"`
	Country   Country
	CountryID uint `json:"country_id"`
	Status    uint `gorm:"default:1" json:"status"`
}

type PaymentMethod struct {
	BaseModel
	Name   string `gorm:"not null;unique" json:"name"`
	Status uint   `gorm:"default:1" json:"status"`
}

type ShippingMethod struct {
	BaseModel
	Name   string `gorm:"not null;unique" json:"name"`
	Status uint   `gorm:"default:1" json:"status"`
}
