package models

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Role struct {
	BaseModel
	Name  string `gorm:"not null" json:"name"`
	Slug  string `gorm:"not null" json:"slug"`
	Users []User `gorm:"foreignKey:RoleID"`
}

type User struct {
	ID uint `gorm:"primary" json:"id"`

	Name            string     `gorm:"not null" json:"name"`
	Email           string     `gorm:"unique;not null" json:"email"`
	Phone           string     `gorm:"unique;not null" json:"phone"`
	EmailVerifiedAt *time.Time `gorm:"type:timestamp;nullable" json:"email_verified_at"`
	Password        string     `gorm:"not null" json:"password"`
	RememberToken   string     `gorm:"size:100" json:"remember_token"`
	RoleID          uint       `gorm:"not null" json:"role_id"`
	Status          uint       `gorm:"default:1" json:"status"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func NewUser(name, email, phone, password string, roleID uint, status uint, emailVerifiedAt *time.Time) *User {
	return &User{
		Name:            name,
		Email:           email,
		Phone:           phone,
		EmailVerifiedAt: emailVerifiedAt,
		Password:        password,
		RoleID:          roleID,
		Status:          status,
	}
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(*gorm.DB) (err error) {
	u.Password, err = HashPassword(u.Password)
	return err
}

func (u *User) BeforeUpdate(*gorm.DB) (err error) {
	u.Password, err = HashPassword(u.Password)
	return err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func ValidToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	return uid == n
}
