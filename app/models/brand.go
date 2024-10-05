package models

type Brand struct {
	BaseModel
	Name               string `validate:"required;unique" json:"name"`
	Description        string `validate:"nullable" json:"description"`
	SeoUrl             string `validate:"nullable" json:"seo_url"`
	SeoMetaTitle       string `validate:"nullable" json:"seo_meta_title"`
	SeoMetaDescription string `validate:"nullable" json:"seo_meta_description"`
	Status             uint   `validate:"default:1" json:"status"`
	Logo               string `validate:"nullable" json:"logo"`
	Banner             string `validate:"nullable" json:"banner"`
}

func NewBrand(name, description, seoMetaTitle, seoMetaDescription string, status uint, logo, banner string, slug string) Brand {
	return Brand{
		Name:               name,
		Description:        description,
		SeoUrl:             slug,
		SeoMetaTitle:       seoMetaTitle,
		SeoMetaDescription: seoMetaDescription,
		Status:             status,
		Logo:               logo,
		Banner:             banner,
	}
}
