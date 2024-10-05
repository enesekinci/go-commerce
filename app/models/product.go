package models

import (
	"time"
)

type ProductCategory struct {
	BaseModel
	ProductID  uint `json:"product_id"`
	CategoryID uint `json:"category_id"`
}

type ProductAttribute struct {
	BaseModel
	ProductID uint `json:"product_id"`
}

type ProductVariant struct {
	BaseModel
	ProductID uint `json:"product_id"`
	VariantID uint `json:"variant_id"`
}

type ProductTag struct {
	BaseModel
	ProductID uint `json:"product_id"`
	TagID     uint `json:"tag_id"`
}

type ProductMedia struct {
	BaseModel
	ProductID uint `json:"product_id"`
	MediaID   uint `json:"media_id"`
}

type Product struct {
	BaseModel
	Name                  string    `gorm:"not null" json:"name"`
	ShortDescription      string    `gorm:"type:text;nullable" json:"short_description"`
	Description           string    `gorm:"type:text;nullable" json:"description"`
	IsVirtual             uint      `gorm:"default:0" json:"is_virtual"`
	Status                uint      `gorm:"default:1" json:"status"`
	Price                 float64   `gorm:"default:0" json:"price"`
	SpecialPrice          uint      `gorm:"nullable" json:"special_price"`
	SpecialPriceType      uint      `gorm:"nullable" json:"special_price_type"`
	SpecialPriceStartDate time.Time `gorm:"nullable" json:"special_price_start_date"`
	SpecialPriceEnd       time.Time `gorm:"nullable" json:"special_price_end" json:"special_price_end"`
	SKU                   string    `gorm:"type:varchar(255);not null;unique" json:"sku"`
	InventoryManagement   uint      `gorm:"default:0" json:"inventory_management"`
	StockAvailability     uint      `gorm:"default:0" json:"stock_availability"`
	SeoURL                string    `gorm:"type:varchar(255);not null;unique" json:"seo_url"`
	SeoMetaTitle          string    `gorm:"type:varchar(255);nullable" json:"seo_meta_title"`
	SeoMetaDescription    string    `gorm:"type:varchar(255);nullable" json:"seo_meta_description"`

	BrandID    uint        `json:"brand_id"`
	Brand      Brand       `gorm:"foreignKey:BrandID" json:"brand"`
	Categories []Category  `gorm:"many2many:product_categories;" json:"categories"`
	TaxID      uint        `json:"tax_id"`
	Tax        Tax         `gorm:"foreignKey:TaxID" json:"tax"`
	Attributes []Attribute `gorm:"many2many:product_attributes;" json:"attributes"`
	Variants   []Variant   `gorm:"many2many:product_variants;" json:"variants"`
	Tags       []Tag       `gorm:"many2many:product_tags;" json:"tags"`
	Medias     []Media     `gorm:"many2many:product_medias;" json:"medias"`
}
