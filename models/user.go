package models

import (
	"dibagi/helpers"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/rs/xid"
	"gorm.io/gorm"
)

type User struct {
	ID        string     `gorm:"primaryKey" json:"id"`
	Email     string     `gorm:"not null;uniqueIndex" json:"email" valid:"required~Email is required,email~Invalid email address"`
	UserName  string     `gorm:"not null;uniqueIndex" json:"user_name" valid:"required~Username is required"`
	Password  string     `gorm:"not null" json:"password" valid:"required~Password is required,minstringlength(8)"`
	FullName  string     `gorm:"not null" json:"full_name" valid:"required~Full Name is required"`
	Age       int        `gorm:"not null" json:"age" valid:"required~Age is required"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type CreateUserResponse struct {
	ID        string     `json:"id"`
	Email     string     `json:"email"`
	UserName  string     `json:"user_name"`
	FullName  string     `json:"full_name"`
	Age       int        `json:"age"`
	CreatedAt *time.Time `json:"created_at"`
}
type EditUserResponse struct {
	ID         string     `json:"id"`
	Email      string     `json:"email"`
	UserName   string     `json:"user_name"`
	FullName   string     `json:"full_name"`
	Age        int        `json:"age"`
	Updated_at *time.Time `json:"updated_at"`
}

type GetUserResponse struct {
	ID         string     `json:"id"`
	Email      string     `json:"email"`
	UserName   string     `json:"user_name"`
	FullName   string     `json:"full_name"`
	Age        int        `json:"age"`
	Created_at *time.Time `json:"created_at"`
	Updated_at *time.Time `json:"updated_at"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}
	newID := xid.New().String()
	u.ID = newID
	hashedPassword, err := helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}

	userName := strings.ReplaceAll(u.UserName, " ", "")
	lowerCase := strings.ToLower(userName)
	u.UserName = lowerCase
	u.Password = hashedPassword
	return nil
}
