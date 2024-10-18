package repository

import (
	"time"

	"github.com/RomanshkVolkov/test-api/internal/core/domain"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	users := []*domain.User{
		{
			UserData: domain.UserData{
				Username: "jose_dwit",
				Email:    "joseguzmandev@gmail.com",
				Name:     "Jose Guzman",
				Role:     domain.Root,
			},
			Password: "password",
		},
		{
			UserData: domain.UserData{
				Username: "diego_dwit",
				Email:    "diegogutcat@gmail.com",
				Name:     "Diego Gutierrez",
				Role:     domain.Root,
			},
			Password: "password",
		},
		{
			UserData: domain.UserData{
				Username: "itzel_dwit",
				Email:    "itzelramonf@gmail.com",
				Name:     "Itzram",
				Role:     domain.Root,
			},
			Password: "password",
		},
	}

	for _, user := range users {
		hashedPassword, _ := HashPassword(user.Password)
		user.Password = hashedPassword

		db.Create(&user)
	}
}

func (database *DSNSource) FindByUsername(username string) (domain.User, error) {
	user := domain.User{}
	database.DB.Model(&domain.User{}).Where("username = ?", username).First(&user)

	if user.ID == 0 {
		return domain.User{}, nil
	}

	return user, nil
}

func (database *DSNSource) FindByUsernameOrEmail(username, email string) (domain.User, error) {
	user := domain.User{}
	database.DB.Model(&domain.User{}).Where("username = ? OR email = ?", username, email).First(&user)

	if user.ID == 0 {
		return domain.User{}, nil
	}

	return user, nil
}

func (database *DSNSource) FindByID(id uint) (domain.User, error) {
	user := domain.User{}
	database.DB.Model(&domain.User{}).Where("id = ?", id).First(&user)

	if user.ID == 0 {
		return domain.User{}, nil
	}

	return user, nil
}

func (database *DSNSource) FindByUsernameAndOTP(username string) (domain.User, error) {
	user := domain.User{}
	if err := database.DB.Model(&domain.User{}).Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (database *DSNSource) FindAndValidateOTP(username string, otp string) (domain.User, map[string][]string, error) {
	schemaError := map[string][]string{}
	user, err := database.FindByUsernameAndOTP(username)
	if err != nil || user.ID == 0 {
		schemaError["username"] = []string{"Usuario no encontrado"}
		return domain.User{}, schemaError, err
	}

	if user.OTP != otp {
		schemaError["code"] = []string{"Código OTP incorrecto"}
		return domain.User{}, schemaError, nil
	}

	if user.OTPExpirationDate.Before(time.Now().UTC()) {
		schemaError["otp"] = []string{"Tu código ha expirado"}
		return domain.User{}, schemaError, nil
	}

	return user, schemaError, nil
}

func (database *DSNSource) NewUser(request *domain.NewUser) (domain.UserData, error) {
	user := domain.User{
		UserData: domain.UserData{
			Username: request.Username,
			Name:     request.Name,
			Email:    request.Email,
			Role:     request.Role,
		},
		Password: request.Password,
	}

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return domain.UserData{}, err
	}
	user.Password = hashedPassword

	if err := database.DB.Create(&user).Error; err != nil {
		return domain.UserData{}, err
	}

	user.Password = MaskString(user.Password)

	return domain.UserData{
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
		Role:     user.Role,
	}, nil
}

func (database *DSNSource) SaveOTPCode(username string) (domain.User, error) {
	user, err := database.FindByUsername(username)
	if err != nil {
		return domain.User{}, err
	}

	if user.ID == 0 {
		return user, nil
	}

	otpCode := GenerateOTP(user.Username)
	user.OTP = otpCode
	user.OTPExpirationDate = time.Now().Add(time.Minute * 1)

	if err := database.DB.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func GenerateOTP(txt string) string {
	base := TxtToRandomNumbers(txt + "otp" + CurrentTime())
	return base[:5]
}

func UserNotFound() domain.APIResponse[string, any] {
	return domain.APIResponse[string, any]{
		Success: false,
		Message: domain.Message{
			En: "User not found",
			Es: "Usuario no encontrado",
		},
	}
}

func (database *DSNSource) UpdatePassword(userID uint, password string) error {
	user, err := database.FindByID(userID)
	if err != nil {
		return err
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	if err := database.DB.Save(&user).Error; err != nil {
		return err
	}

	return nil
}
