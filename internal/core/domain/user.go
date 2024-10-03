package domain

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	Root     Role = "root"
	Admin    Role = "admin"
	Customer Role = "customer"
	Employer Role = "employer"
)

type UserData struct {
	Username string `gorm:"type nvarchar(200);not null;unique;" json:"username" validate:"required,min=6,max=200"`
	Name     string `gorm:"type nvarchar(300);not null" json:"name" validate:"required,min=3,max=300"`
	Email    string `gorm:"type nvarchar(300);not null;unique;" json:"email" validate:"required,email,max=300"`
	Role     Role   `gorm:"type nvarchar(50);not null" json:"role" validate:"required"`
}

type User struct {
	gorm.Model
	UserData
	OTP               string    `gorm:"type nvarchar(5)" json:"otp"` // One Time Password
	OTPExpirationDate time.Time `gorm:"column otp_expiration_date" json:"otpExpirationDate"`
	Password          string    `gorm:"type nvarchar(200);not null" json:"-" validate:"required,min=6,max=200"`
}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PasswordResetRequest struct {
	Username string `json:"username,omitempty"`
}

type NewUser struct {
	UserData
	Password string `json:"password" validate:"required,min=6"`
}

type ForgottenPasswordCode struct {
	Username string `json:"username" validate:"required,min=6,max=200"`
	OTP      string `json:"otp" validate:"required,min=5,max=5"`
}

type ResetForgottenPassword struct {
	Username        string `json:"username" validate:"required,min=6,max=200"`
	OTP             string `json:"otp" validate:"required,min=5,max=5"`
	Password        string `json:"password" validate:"required,min=6,max=200"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=6,max=200"`
}

type ChangePassword struct {
	CurrentPassword string `json:"currentPassword" validate:"required,min=6,max=200"`
	Password        string `json:"password" validate:"required,min=6,max=200"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=6,max=200"`
}
