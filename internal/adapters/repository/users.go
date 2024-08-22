package repository

import (
	"fmt"
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
			Password: "root",
		},
		{
			UserData: domain.UserData{
				Username: "diego_dwit",
				Email:    "diegogutcat@gmail.com",
				Name:     "Diego Gutierrez",
				Role:     domain.Root,
			},
			Password: "root",
		}, {
			UserData: domain.UserData{
				Username: "itzel_dwit",
				Email:    "itzram@gmail.com",
				Name:     "Itzram",
				Role:     domain.Root,
			},
			Password: "root",
		},
	}

	for _, user := range users {
		hashedPassword, _ := HashPassword(user.Password)
		user.Password = hashedPassword

		db.Create(&user)
	}
}

func FindByUsername(username string) (domain.User, error) {
	user := domain.User{}
	DBSQLServer.Model(&domain.User{}).Where("username = ?", username).First(&user)
	fmt.Println(user)
	if user.ID == 0 {
		return domain.User{}, nil
	}

	return user, nil
}

func FindByUsernameOrEmail(username, email string) (domain.User, error) {
	user := domain.User{}
	DBSQLServer.Model(&domain.User{}).Where("username = ? OR email = ?", username, email).First(&user)
	fmt.Println(user)
	if user.ID == 0 {
		return domain.User{}, nil
	}

	return user, nil
}

func FindByID(id uint) (domain.User, error) {
	user := domain.User{}
	DBSQLServer.Model(&domain.User{}).Where("id = ?", id).First(&user)
	fmt.Println(user)
	if user.ID == 0 {
		return domain.User{}, nil
	}

	return user, nil
}

func FindByUsernameAndOTP(username string) (domain.User, error) {
	user := domain.User{}
	if err := DBSQLServer.Model(&domain.User{}).Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func FindAndValidateOTP(username string, otp string) (domain.User, map[string][]string, error) {
	schemaError := map[string][]string{}
	user, err := FindByUsernameAndOTP(username)
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

func NewUser(request *domain.NewUser) (domain.UserData, error) {
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

	if err := DBSQLServer.Create(&user).Error; err != nil {
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

func SaveOTPCode(username string) (domain.User, error) {
	user, err := FindByUsername(username)
	if err != nil {
		return domain.User{}, err
	}

	if user.ID == 0 {
		return user, nil
	}

	otpCode := GenerateOTP(user.Username)
	user.OTP = otpCode
	user.OTPExpirationDate = time.Now().Add(time.Minute * 1)

	if err := DBSQLServer.Save(&user).Error; err != nil {
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
		Message: "User not found",
	}
}

func UpdatePassword(userID uint, password string) error {
	user, err := FindByID(userID)
	if err != nil {
		return err
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	if err := DBSQLServer.Save(&user).Error; err != nil {
		return err
	}

	return nil
}
