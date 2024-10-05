package models

type Tag struct {
	BaseModel
	Name string `validate:"required;unique" json:"name"`
}

func NewTag(name string) *Tag {
	return &Tag{
		Name: name,
	}
}
