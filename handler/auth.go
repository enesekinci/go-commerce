package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go-commerce/app/models"
	"go-commerce/app/resource"
	"go-commerce/config"
	"go-commerce/core/constant"
	"go-commerce/core/helper"
	"go-commerce/database"
	"gorm.io/gorm"
	"time"
)

func getUserByEmail(e string) (*models.User, error) {
	db := database.DB
	var user models.User
	if err := db.Where(&models.User{Email: e}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func getUserByPhone(p string) (*models.User, error) {
	db := database.DB
	var user models.User
	if err := db.Where(&models.User{Phone: p}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func RefreshToken(c *fiber.Ctx) error {

	type LoginInput struct {
		Email    string `validate:"required,email" json:"email"`
		Password string `validate:"required,min=6,alphanum" json:"password"`
	}

	input := new(LoginInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "Review your request", "error": err})
	}

	result := helper.ValidateStruct(input)

	if result != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "Review your request", "errors": result})
	}

	email := input.Email
	password := input.Password
	user, err := new(models.User), *new(error)

	user, err = getUserByEmail(email)

	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": false, "message": "User not found", "errors": map[string]interface{}{"user": constant.NotFound}})
	}

	if !helper.CheckPasswordHash(password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": false, "message": "Invalid password"})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["role_id"] = user.RoleID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24 * 60).Unix()
	claims["iat"] = time.Now().Unix()
	claims["token_type"] = "refresh"

	t, err := token.SignedString([]byte(config.Config("JWT_REFRESH_SECRET_KEY")))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"status":        true,
		"message":       "Success login",
		"refresh_token": t,
		"exp":           claims["exp"],
		"user":          resource.NewUserResource(user),
	})
}

func AccessToken(c *fiber.Ctx) error {
	refreshToken := helper.GetToken(c)

	println("refreshToken", refreshToken)

	tokenClaims, _ := helper.GetTokenClaims(refreshToken, true)

	println("tokenClaims", tokenClaims)

	if tokenClaims == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": false, "message": "Unauthorized"})
	}

	user, err := getUserByEmail(tokenClaims["email"].(string))

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": false, "message": "Unauthorized"})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["role_id"] = user.RoleID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["iat"] = time.Now().Unix()
	claims["token_type"] = "access"

	t, err := token.SignedString([]byte(config.Config("JWT_SECRET_KEY")))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": true,
		"message":      "Success login",
		"access_token": t,
		"exp":          claims["exp"],
		"user":         resource.NewUserResource(user),
	})
}

func Register(c *fiber.Ctx) error {

	type RegisterInput struct {
		Name     string `validate:"required" json:"name"`
		Email    string `validate:"required,email" json:"email"`
		Phone    string `validate:"required" json:"phone"`
		Password string `validate:"required" json:"password"`
	}

	input := new(RegisterInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "Review your request", "error": err})
	}

	result := helper.ValidateStruct(input)

	if result != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "Review your request", "errors": result})
	}

	if user, _ := getUserByEmail(input.Email); user.Email != "" {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": false, "message": "User already exists", "errors": map[string]interface{}{"email": constant.AlreadyTaken}})
	}

	if user, _ := getUserByPhone(input.Phone); user.Phone != "" {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": false, "message": "User already exists", "errors": map[string]interface{}{"phone": constant.AlreadyTaken}})
	}

	user := models.NewUser(input.Name, input.Email, input.Phone, input.Password, 3, 1, nil)

	result = database.DB.Create(&user)

	println(result)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": true, "message": "User successfully created", "user": user})
}
func ForgotPassword(c *fiber.Ctx) error {
	type ForgotPasswordInput struct {
		Email string `validate:"required,email" json:"email"`
	}

	input := new(ForgotPasswordInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "Review your request", "error": err})
	}

	result := helper.ValidateStruct(input)

	if result != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "Review your request", "errors": result})
	}

	user, err := getUserByEmail(input.Email)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": false, "message": "Internal server error"})
	}

	if user == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": false, "message": "User not found"})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	claims["iat"] = time.Now().Unix()
	claims["token_type"] = "forgot"

	t, err := token.SignedString([]byte(config.Config("JWT_SECRET_KEY")))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	body := "Click this link to reset your password: " + config.Config("APP_URL") + "/auth/password/change?token=" + t

	err = helper.SendMail(user.Email, "Forgot Password", body)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": true, "message": "Success create forgot password token"})
}
func ChangePassword(c *fiber.Ctx) error {
	type ChangePasswordInput struct {
		CurrentPassword string `validate:"required" json:"current_password"`
		Password        string `validate:"required" json:"password"`
	}

	input := new(ChangePasswordInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "Review your request", "error": err})
	}

	result := helper.ValidateStruct(input)

	if result != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "Review your request", "errors": result})
	}

	token := helper.GetToken(c)

	claims, err := helper.GetTokenClaims(token, false)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": false, "message": "Unauthorized"})
	}

	user, err := getUserByEmail(claims["email"].(string))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": false, "message": "Internal server error"})
	}

	if user == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": false, "message": "User not found"})
	}

	if !helper.CheckPasswordHash(input.CurrentPassword, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": false, "message": "Invalid password"})
	}

	user.Password = helper.HashPassword(input.Password)

	database.DB.Save(&user)

	return c.JSON(fiber.Map{"status": true, "message": "Success change password"})
}
func UpdateProfile(c *fiber.Ctx) error {
	type UpdateProfileInput struct {
		Name  string `validate:"required" json:"name"`
		Phone string `validate:"required" json:"phone"`
	}

	input := new(UpdateProfileInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "Review your request", "error": err})
	}

	result := helper.ValidateStruct(input)

	if result != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": "Review your request", "errors": result})
	}

	token := helper.GetToken(c)

	claims, err := helper.GetTokenClaims(token, false)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": false, "message": "Unauthorized"})
	}

	user, err := getUserByEmail(claims["email"].(string))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": false, "message": "Internal server error"})
	}

	if user == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": false, "message": "User not found"})
	}

	user.Name = input.Name
	user.Phone = input.Phone

	database.DB.Save(&user)

	return c.JSON(fiber.Map{"status": true, "message": "Success update profile", "user": user})
}
