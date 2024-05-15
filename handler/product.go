package handler

import (
	"github.com/gofiber/fiber/v2"
	"go-commerce/app/models"
	"go-commerce/core/helper"
	"go-commerce/database"
	"time"
)

//GetProducts
//CreateProduct
//GetProduct
//UpdateProduct
//DeleteProduct

type CreateProductInput struct {
	Name                  string    `json:"name" validate:"required,type:text,max=255"`
	ShortDescription      string    `json:"short_description" validate:"nullable;type:text"`
	Description           string    `json:"description" validate:"nullable;type:text"`
	IsVirtual             uint      `json:"is_virtual" validate:"default:0"`
	Status                uint      `json:"status" validate:"default:1"`
	Price                 float64   `json:"price" validate:"default:0"`
	SpecialPrice          uint      `json:"special_price" validate:"nullable"`
	SpecialPriceType      uint      `json:"special_price_type" validate:"nullable"`
	SpecialPriceStartDate time.Time `json:"special_price_start_date" validate:"nullable"`
	SpecialPriceEnd       time.Time `json:"special_price_end" validate:"nullable"`
	SKU                   string    `json:"sku" validate:"type:varchar(255);not null;unique"`
	InventoryManagement   uint      `json:"inventory_management" validate:"default:0"`
	StockAvailability     uint      `json:"stock_availability" validate:"default:0"`
	SeoURL                string    `json:"seo_url" validate:"type:varchar(255);not null;unique"`
	SeoMetaTitle          string    `json:"seo_meta_title" validate:"type:varchar(255);nullable"`
	SeoMetaDescription    string    `json:"seo_meta_description" validate:"type:varchar(255);nullable"`
	BrandID               uint      `json:"brand_id"`
	Categories            []uint    `json:"categories"`
	TaxID                 uint      `json:"tax_id"`
	Attributes            []uint    `json:"attributes"`
	Variants              []uint    `json:"variants"`
	Tags                  []uint    `json:"tags"`
	Medias                []uint    `json:"medias"`
}

func GetProducts(c *fiber.Ctx) error {
	var products []models.Product

	database.DB.Find(&products)

	return c.JSON(fiber.Map{"status": "success", "message": "Products found", "products": products})
}
func CreateProduct(c *fiber.Ctx) error {
	input := new(CreateProductInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Review your request", "error": err})
	}

	result := helper.ValidateStruct(input)

	if result != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Review your request", "errors": result})
	}

	if database.IsExistInDB("products", "name", input.Name) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Product already exist", "data": nil})
	}

	product := models.Product{
		Name:                  input.Name,
		ShortDescription:      input.ShortDescription,
		Description:           input.Description,
		IsVirtual:             input.IsVirtual,
		Status:                input.Status,
		Price:                 input.Price,
		SpecialPrice:          input.SpecialPrice,
		SpecialPriceType:      input.SpecialPriceType,
		SpecialPriceStartDate: input.SpecialPriceStartDate,
		SpecialPriceEnd:       input.SpecialPriceEnd,
		SKU:                   input.SKU,
		InventoryManagement:   input.InventoryManagement,
		StockAvailability:     input.StockAvailability,
		SeoURL:                input.SeoURL,
		SeoMetaTitle:          input.SeoMetaTitle,
		SeoMetaDescription:    input.SeoMetaDescription,
		BrandID:               input.BrandID,
		TaxID:                 input.TaxID,
	}

	database.DB.Create(&product)

	return c.JSON(fiber.Map{"status": "success", "message": "Product created", "product": product})
}
func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	var product models.Product

	database.DB.Find(&product, id)

	if product.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Product found", "product": product})
}
func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	input := new(CreateProductInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Review your request", "error": err})
	}

	result := helper.ValidateStruct(input)

	if result != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Review your request", "errors": result})
	}

	var product models.Product

	database.DB.Find(&product, id)

	if product.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})

	}

	product.Name = input.Name
	product.ShortDescription = input.ShortDescription
	product.Description = input.Description
	product.IsVirtual = input.IsVirtual
	product.Status = input.Status
	product.Price = input.Price
	product.SpecialPrice = input.SpecialPrice
	product.SpecialPriceType = input.SpecialPriceType
	product.SpecialPriceStartDate = input.SpecialPriceStartDate
	product.SpecialPriceEnd = input.SpecialPriceEnd
	product.SKU = input.SKU
	product.InventoryManagement = input.InventoryManagement
	product.StockAvailability = input.StockAvailability
	product.SeoURL = input.SeoURL
	product.SeoMetaTitle = input.SeoMetaTitle
	product.SeoMetaDescription = input.SeoMetaDescription
	product.BrandID = input.BrandID
	product.TaxID = input.TaxID

	database.DB.Save(&product)

	return c.JSON(fiber.Map{"status": "success", "message": "Product updated", "product": product})
}
func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	var product models.Product

	database.DB.Find(&product, id)

	if product.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})
	}

	database.DB.Delete(&product)

	return c.JSON(fiber.Map{"status": "success", "message": "Product deleted", "product": product})
}
