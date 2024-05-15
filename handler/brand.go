package handler

import (
	"github.com/gofiber/fiber/v2"
	"go-commerce/app/models"
	"go-commerce/core/helper"
	"go-commerce/database"
)

type CreateBrandInput struct {
	Name               string `validate:"required;" json:"name"`
	Description        string `validate:"nullable" json:"description"`
	SeoMetaTitle       string `validate:"nullable" json:"seo_meta_title"`
	SeoMetaDescription string `validate:"nullable" json:"seo_meta_description"`
	Status             uint   `validate:"default:1" json:"status"`
	Logo               string `validate:"nullable" json:"logo"`
	Banner             string `validate:"nullable" json:"banner"`
}

func GetBrands(c *fiber.Ctx) error {
	var brands []models.Brand

	database.DB.Find(&brands)

	return c.JSON(fiber.Map{"status": "success", "message": "Categories found", "brands": brands})
}

func GetBrand(c *fiber.Ctx) error {

	id := c.Params("id")

	var brand models.Brand

	database.DB.Find(&brand, id)

	if brand.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No category found with ID", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Category found", "brand": brand})
}

func CreateBrand(c *fiber.Ctx) error {
	input := new(CreateBrandInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Review your request", "error": err})
	}

	result := helper.ValidateStruct(input)

	if result != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Review your request", "errors": result})
	}

	slug := helper.NewSlug(input.Name)

	for !database.IsExistInDB("brands", "seo_url", slug) {
		slug = helper.IncrementSlug(slug)
	}

	brand := models.NewBrand(input.Name, input.Description, input.SeoMetaTitle, input.SeoMetaDescription, input.Status, input.Logo, input.Banner, slug)

	database.DB.Create(&brand)

	return c.JSON(fiber.Map{"status": "success", "message": "Brand created", "brand": brand})
}

func UpdateBrand(c *fiber.Ctx) error {
	id := c.Params("id")

	var brand models.Brand

	database.DB.Find(&brand, id)

	if brand.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No category found with ID", "data": nil})
	}

	input := new(CreateBrandInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Review your request", "error": err})
	}

	result := helper.ValidateStruct(input)

	if result != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Review your request", "errors": result})
	}

	if brand.Name != input.Name {
		if database.IsExistInDB("brands", "name", input.Name) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Brand already exist", "data": nil})
		}

		brand.SeoUrl = helper.NewSlug(input.Name)

		for !database.IsExistInDB("brands", "seo_url", brand.SeoUrl) {
			brand.SeoUrl = helper.IncrementSlug(brand.SeoUrl)
		}
	}

	brand.Name = input.Name
	brand.Description = input.Description
	brand.SeoMetaTitle = input.SeoMetaTitle
	brand.SeoMetaDescription = input.SeoMetaDescription
	brand.Status = input.Status
	brand.Logo = input.Logo
	brand.Banner = input.Banner

	database.DB.Save(&brand)

	return c.JSON(fiber.Map{"status": "success", "message": "Brand updated", "brand": brand})
}

func DeleteBrand(c *fiber.Ctx) error {
	id := c.Params("id")

	var brand models.Brand

	database.DB.Find(&brand, id)

	if brand.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No category found with ID", "data": nil})
	}

	database.DB.Delete(&brand)

	return c.JSON(fiber.Map{"status": "success", "message": "Brand deleted", "brand": brand})
}
