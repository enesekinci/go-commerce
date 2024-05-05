package resource

import (
	"go-commerce/app/models"
	"go-commerce/core/helper"
)

type UserResource struct {
	Id              int    `json:"id"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Name            string `json:"name"`
	EmailVerifiedAt string `json:"email_verified_at"`
	RoleID          int    `json:"role_id"`
	Status          int    `json:"status"`
}

func NewUserResource(user *models.User) *UserResource {
	return &UserResource{
		Id:              int(user.ID),
		Email:           user.Email,
		Phone:           user.Phone,
		Name:            user.Name,
		EmailVerifiedAt: helper.FormatTime(user.EmailVerifiedAt),
		RoleID:          int(user.RoleID),
		Status:          int(user.Status),
	}
}
