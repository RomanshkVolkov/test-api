package domain

import (
	"gorm.io/gorm"
)

type Dev struct {
	gorm.Model
	Tag string `gorm:"type nvarchar(200);not null" json:"tag" validate:"required,min=3,max=200"`
	IP  string `gorm:"type nvarchar(200);not null" json:"ip" validate:"required,min=3,max=200"`
}
