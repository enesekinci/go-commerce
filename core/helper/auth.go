package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go-commerce/config"
)

func GetTokenClaims(tokenString string, isRefresh bool) (jwt.MapClaims, error) {

	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		tokenKey := "JWT_SECRET_KEY"

		if isRefresh {
			tokenKey = "JWT_REFRESH_SECRET_KEY"
		}
		return []byte(config.Config(tokenKey)), nil
	})

	if err != nil {
		return nil, err
	}

	for key, value := range claims {
		println("Key:", key, "Value:", value)
	}

	return claims, err
}

func GetToken(c *fiber.Ctx) string {

	token := c.Get("Authorization")

	token = token[7:]

	return token
}
