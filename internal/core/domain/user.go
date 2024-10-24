package domain

type UserData struct {
	Username  string       `gorm:"type nvarchar(200);not null;unique;" json:"username" validate:"required,min=6,max=200"`
	Name      string       `gorm:"type nvarchar(300);not null" json:"name" validate:"required,min=3,max=300"`
	Email     string       `gorm:"type nvarchar(300);not null;unique;" json:"email" validate:"required,email,max=300"`
	ProfileID uint         `gorm:"not null" json:"-"`
	Profile   UserProfiles `gorm:"foreignKey:ProfileID;references:ID" json:"profile"`
}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInResponse struct {
	UserData
	Token string `json:"token"`
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
	OTP      string `json:"otp" validate:"required,min=6,max=6"`
}

type ResetForgottenPassword struct {
	Username        string `json:"username" validate:"required,min=6,max=200"`
	OTP             string `json:"otp" validate:"required,min=6,max=6"`
	Password        string `json:"password" validate:"required,min=6,max=200"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=6,max=200"`
}

type ChangePassword struct {
	CurrentPassword string `json:"currentPassword" validate:"required,min=6,max=200"`
	Password        string `json:"password" validate:"required,min=6,max=200"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=6,max=200"`
}
