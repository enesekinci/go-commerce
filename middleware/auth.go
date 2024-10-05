package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go-commerce/config"
	"go-commerce/core/constant"
)

func VerifyAccessToken() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(config.Config("JWT_SECRET_KEY"))},
		ErrorHandler: jwtError,
		SuccessHandler: func(c *fiber.Ctx) error {
			return c.Next()
		},
	})
}

func VerifyRefreshToken() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(config.Config("JWT_REFRESH_SECRET_KEY"))},
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	println(err.Error())
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": false, "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": false, "message": "Invalid or expired JWT", "data": nil})
}

func IsAdmin(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)

	claims := user.Claims.(jwt.MapClaims)

	roleId := uint(claims["role_id"].(float64))

	allowedRoles := []uint{
		constant.SuperAdmin,
		constant.Admin,
	}

	isExist := false

	for _, role := range allowedRoles {
		if role == roleId {
			isExist = true
			break
		}
	}

	if !isExist {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": false, "message": "Unauthorized", "data": nil})
	}

	return c.Next()
}
