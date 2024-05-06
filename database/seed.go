package database

import (
	"go-commerce/app/models"
	"go-commerce/core/constant"
	"gorm.io/gorm"
)

func DropTables(db *gorm.DB) {
	_ = db.Migrator().DropTable(
		&models.User{},
		&models.Role{},
		&models.Category{},
		&models.Brand{},

		&models.Tax{},
		&models.TaxClass{},
		&models.Setting{},
		&models.Media{},
		&models.Language{},
		&models.Currency{},
		&models.Country{},
		&models.Zone{},
		&models.PaymentMethod{},
		&models.ShippingMethod{},

		&models.Attribute{},
		&models.OptionType{},
		&models.Option{},
		&models.Tag{},
		&models.Variant{},
		&models.Product{},
		&models.ProductVariant{},
	)
}
func SeedData(db *gorm.DB) {

	roles := []models.Role{
		{
			Name: "Super Admin",
			Slug: "super-admin",
		},
		{
			Name: "Admin",
			Slug: "admin",
		},
		{
			Name: "User",
			Slug: "user",
		},
	}

	users := []models.User{
		{
			Name:     "Enes Ekinci",
			Email:    "enes.eknc.96@gmail.com",
			Phone:    "5369501299",
			Password: "password",
			RoleID:   constant.SuperAdmin,
			Status:   constant.ACTIVE,
		},
		{
			Name:     "John Doe",
			Email:    "john@gmail.com",
			Phone:    "5555555555",
			Password: "password",
			RoleID:   constant.User,
			Status:   constant.ACTIVE,
		},
		{
			Name:     "Alex Doe",
			Email:    "alex@gmail.com",
			Phone:    "5545545454",
			Password: "password",
			RoleID:   constant.Admin,
			Status:   constant.INACTIVE,
		},
	}

	optionTypes := []string{
		"Field",
		"Textarea",
		"Dropdown",
		"Checkbox",
		"Radio",
		"Multiselect",
		"Date",
		"DateTime",
		"Time",
	}

	for _, role := range roles {
		db.Create(&role)
	}

	for _, user := range users {
		db.Create(&user)
	}

	for _, optionType := range optionTypes {
		db.Create(&models.OptionType{Name: optionType})
	}
}
