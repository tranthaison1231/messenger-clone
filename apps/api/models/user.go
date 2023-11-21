package models

import (
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primary_key;autoIncrement"`
	Email    string `gorm:"uniqueIndex;not null;type:varchar(255)"`
	Salt     string `gorm:"not null"`
	Avatar   string `gorm:"type:text"`
	Password string `gorm:"not null"`
	Gender   string `gorm:"not null"`
	// Chats     []*Chat   `gorm:"many2many:chat_users;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserResponse struct {
	ID        uint      `json:"id,omitempty"`
	Email     string    `json:"email,omitempty"`
	Gender    string    `json:"gender,omitempty"`
	Avatar    string    `json:"avatar,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Gender string

type SignUpRequest struct {
	SignInRequest
	Gender string `json:"gender" validate:"required"`
	Avatar string `json:"avatar"`
}

func (u *User) ValidatePwdStaticHash(password string) error {
	if password == "" {
		return errors.WithStack(errors.New("Password cannot be empty"))
	}
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}