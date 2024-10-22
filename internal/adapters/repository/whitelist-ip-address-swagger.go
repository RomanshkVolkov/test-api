package repository

import (
	"github.com/RomanshkVolkov/test-api/internal/core/domain"
)

func (database *DSNSource) SwaggerValidateIPAddress(ip string) (bool, error) {
	dev := []domain.Dev{}
	if err := database.DB.Model(&domain.Dev{}).Where("ip = ?", ip).First(&dev).Error; err != nil {
		return false, err
	}

	if len(dev) == 0 {
		return false, nil
	}

	return true, nil

}
