package resource

import "go-commerce/app/models"

type CategoryResource struct {
	ID                 uint   `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Parent             uint   `json:"parent"`
	SeoUrl             string `json:"seo_url"`
	SeoMetaTitle       string `json:"seo_meta_title"`
	SeoMetaDescription string `json:"seo_meta_description"`
	Searchable         int    `json:"searchable"`
	Status             int    `json:"status"`
	Logo               string `json:"logo"`
	Banner             string `json:"banner"`
}

func NewCategoryResource(category models.Category) *CategoryResource {
	return &CategoryResource{
		ID:                 category.ID,
		Name:               category.Name,
		Description:        category.Description,
		Parent:             category.Parent,
		SeoUrl:             category.SeoUrl,
		SeoMetaTitle:       category.SeoMetaTitle,
		SeoMetaDescription: category.SeoMetaDescription,
		Searchable:         category.Searchable,
		Status:             category.Status,
		Logo:               category.Logo,
		Banner:             category.Banner,
	}
}
