package handler

import (
	"github.com/gofiber/fiber/v2"
	"go-commerce/app/models"
	"go-commerce/app/resource"
	"go-commerce/core/helper"
	"go-commerce/database"
)

type CreateCategoryInput struct {
	Name        string `validate:"required,min=3,max=100" json:"name"`
	Description string `validate:"nullable" json:"description"`
	Parent      uint   `validate:"default:0" json:"parent"`
	//SeoUrl             string `validate:"nullable" json:"seo_url"`
	SeoMetaTitle       string `validate:"nullable" json:"seo_meta_title"`
	SeoMetaDescription string `validate:"nullable" json:"seo_meta_description"`
	Searchable         int    `validate:"default:1;in=0,1" json:"searchable"`
	Status             int    `validate:"default:1;in=0,1" json:"status"`
	Logo               string `validate:"nullable" json:"logo"`
	Banner             string `validate:"nullable" json:"banner"`
}

func GetCategories(c *fiber.Ctx) error {
	var categories []models.Category

	database.DB.Find(&categories)

	return c.JSON(fiber.Map{"status": "success", "message": "Categories found", "categories": categories})
}

func GetCategory(c *fiber.Ctx) error {

	id := c.Params("id")

	var category models.Category

	database.DB.Find(&category, id)

	if category.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No category found with ID", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Category found", "category": resource.NewCategoryResource(category)})
}

func CreateCategory(c *fiber.Ctx) error {
	input := new(CreateCategoryInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Review your request", "error": err})
	}

	result := helper.ValidateStruct(input)

	if result != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Review your request", "errors": result})
	}

	if helper.IsExistInDB("categories", "name", input.Name) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Category already exist", "data": nil})
	}

	category := models.NewCategory(input.Name, input.Description, input.Parent, input.SeoMetaTitle, input.SeoMetaDescription, input.Searchable, input.Status, input.Logo, input.Banner)

	database.DB.Create(&category)

	return c.JSON(fiber.Map{"status": "success", "message": "Category created", "category": category})
}

func UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	input := new(CreateCategoryInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Review your request", "error": err})
	}

	result := helper.ValidateStruct(input)

	if result != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Review your request", "errors": result})
	}

	var category models.Category

	database.DB.Find(&category, id)

	if category.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No category found with ID", "data": nil})
	}

	if category.Name != input.Name {
		if helper.IsExistInDB("categories", "name", input.Name) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Category already exist", "data": nil})
		}

		category.SeoUrl = helper.NewSlug(input.Name, "categories", "seo_url")
	}

	category.Name = input.Name
	category.Description = input.Description
	category.Parent = input.Parent
	category.SeoMetaTitle = input.SeoMetaTitle
	category.SeoMetaDescription = input.SeoMetaDescription
	category.Searchable = input.Searchable
	category.Status = input.Status
	category.Logo = input.Logo
	category.Banner = input.Banner

	database.DB.Save(&category)

	return c.JSON(fiber.Map{"status": "success", "message": "Category updated", "category": category})
}

func DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	var category models.Category

	database.DB.Find(&category, id)

	if category.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No category found with ID", "data": nil})
	}

	database.DB.Delete(&category)

	return c.JSON(fiber.Map{"status": "success", "message": "Category deleted", "data": nil})
}
