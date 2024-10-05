package database

import (
	"fmt"
	"go-commerce/app/models"
	"go-commerce/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	p := config.Config("DB_PORT")

	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		panic("failed to parse database port")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	//DB := &gorm.DB{}
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	DropTables()
	//
	err = DB.AutoMigrate(
		&models.Role{},
		&models.User{},

		&models.Category{},
		&models.Brand{},

		//entities
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
		//entities

		&models.VariantType{},
		&models.Variant{},

		&models.Tag{},

		&models.Attribute{},
		&models.AttributeCategory{},

		&models.OptionType{},
		&models.Option{},

		&models.ProductCategory{},
		&models.ProductAttribute{},
		&models.ProductVariant{},
		&models.ProductTag{},
		&models.ProductMedia{},
		&models.Product{},
	)

	if err != nil {
		panic("failed to migrate database")
	}

	SeedData()

	fmt.Println("Database Migrated")

}
