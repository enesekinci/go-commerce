package helper

import (
	"go-commerce/app/models"
	"go-commerce/database"
	"golang.org/x/crypto/bcrypt"
)

func ValidUser(id string, password string) bool {

	var user models.User

	database.DB.First(&user, id)

	if user.Email == "" {
		return false
	}

	if !CheckPasswordHash(password, user.Password) {
		return false
	}

	return true
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func IsUserExist(email string) bool {

	var user models.User

	database.DB.Where("email = ?", email).First(&user)

	if user.Email == "" {
		return false
	}

	return true
}

func IsUserExistByID(id string) bool {

	var user models.User

	database.DB.First(&user, id)

	if user.Email == "" {
		return false
	}

	return true
}

func IsUserExistByPhone(phone string) bool {

	var user models.User

	database.DB.Where("phone = ?", phone).First(&user)

	if user.Email == "" {
		return false
	}

	return true
}

func IsUserExistByEmailOrPhone(email string, phone string) interface{} {

	var user models.User

	database.DB.Where("email = ? OR phone = ?", email, phone).First(&user)

	if user.Email == "" {
		return false
	}

	return user
}
