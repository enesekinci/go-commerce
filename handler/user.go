package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go-commerce/app/models"
	"go-commerce/app/resource"
	"go-commerce/core/constant"
	"go-commerce/core/helper"
	"go-commerce/database"
)

type CreateUserInput struct {
	Name     string `validate:"required,min=3,max=100" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Phone    string `validate:"required,min=10,max=15,number" json:"phone"`
	Password string `validate:"required,min=6,max=100" json:"password"`
}
type UpdateUserInput struct {
	Name string `validate:"min=3,max=100" json:"name"`
}

type PasswordInput struct {
	Password string `json:"password"`
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User

	database.DB.Find(&user, id)

	if user.Email == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": false, "message": "No user found with ID", "data": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": "User found", "user": resource.NewUserResource(&user)})
}

func CreateUser(c *fiber.Ctx) error {

	input := new(CreateUserInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": false, "message": "Review your request", "error": err})
	}

	result := helper.ValidateStruct(input)

	if result != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "Review your request", "errors": result})
	}

	user := models.NewUser(input.Name, input.Email, input.Phone, input.Password, 1, 1, nil)

	if helper.IsUserExist(user.Email) {
		errors := map[string]interface{}{"email": constant.AlreadyTaken}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "Email already exist", "data": nil, "errors": errors})
	}

	if helper.IsUserExistByPhone(user.Phone) {
		errors := map[string]interface{}{"phone": constant.AlreadyTaken}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "Phone already exist", "data": nil, "errors": errors})
	}

	/*if existUser := helper.IsUserExistByEmailOrPhone(user.Email, user.Phone); existUser != false {
		if existUser.(models.User).Email == user.Email {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "Email already exist", "data": nil})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "Phone already exist", "data": nil})
	}*/

	database.DB.Create(&user)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": true, "message": "User successfully created", "user": resource.NewUserResource(user)})
}

func UpdateUser(c *fiber.Ctx) error {

	input := new(UpdateUserInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": false, "message": "Review your request", "error": err})
	}

	result := helper.ValidateStruct(input)

	if result != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "Review your request", "errors": result})
	}

	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !models.ValidToken(token, id) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": false, "message": "Invalid token id", "data": nil})
	}

	var user models.User

	database.DB.First(&user, id)
	user.Name = input.Name
	database.DB.Save(&user)

	return c.JSON(fiber.Map{"status": true, "message": "User successfully updated", "user": resource.NewUserResource(&user)})
}

func DeleteUserYourself(c *fiber.Ctx) error {

	var input PasswordInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": false, "message": "Review your request", "error": err})
	}

	id := c.Params("id")

	token := c.Locals("user").(*jwt.Token)

	if !models.ValidToken(token, id) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": false, "message": "Invalid token id", "data": nil})
	}

	var user models.User

	database.DB.First(&user, id)

	if !helper.CheckPasswordHash(input.Password, user.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "Password not match", "data": nil})
	}

	database.DB.Delete(&user)

	return c.JSON(fiber.Map{"status": true, "message": "User successfully deleted", "data": nil})
}

func DeleteUserByAdmin(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User

	database.DB.First(&user, id)

	database.DB.Delete(&user)

	return c.JSON(fiber.Map{"status": true, "message": "User successfully deleted", "data": nil})
}
