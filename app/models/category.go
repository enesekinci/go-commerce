package models

type Category struct {
	BaseModel
	Name               string `gorm:"not null;unique" json:"name"`
	Description        string `gorm:"nullable" json:"description"`
	Parent             uint   `gorm:"default:0" json:"parent"`
	SeoUrl             string `gorm:"not null;unique" json:"seo_url"`
	SeoMetaTitle       string `gorm:"nullable" json:"seo_meta_title"`
	SeoMetaDescription string `gorm:"nullable" json:"seo_meta_description"`
	Searchable         int    `gorm:"default:1" json:"searchable"`
	Status             int    `gorm:"default:1" json:"status"`
	Logo               string `gorm:"nullable" json:"logo"`
	Banner             string `gorm:"nullable" json:"banner"`
}

func NewCategory(name string, description string, parent uint, seoMetaTitle string, seoMetaDescription string, searchable int, status int, logo string, banner, slug string) *Category {
	return &Category{
		Name:               name,
		Description:        description,
		Parent:             parent,
		SeoUrl:             slug,
		SeoMetaTitle:       seoMetaTitle,
		SeoMetaDescription: seoMetaDescription,
		Searchable:         searchable,
		Status:             status,
		Logo:               logo,
		Banner:             banner,
	}
}
