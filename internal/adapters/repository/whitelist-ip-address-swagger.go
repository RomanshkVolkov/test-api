package repository

import (
	"github.com/RomanshkVolkov/test-api/internal/core/domain"
)

func SwaggerValidateIPAddress(IP string) (bool, error) {
	dev := []domain.Dev{}
	if err := DBSQLServer.Model(&domain.Dev{}).Where("ip = ?", IP).First(&dev).Error; err != nil {
		return false, err
	}

	if len(dev) == 0 {
		return false, nil
	}

	return true, nil

}
