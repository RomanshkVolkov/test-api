package domain

type Hotel struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"type nvarchar(255);not null" json:"name" validate:"required,min=3,max=255"`
	CompanyName string `gorm:"type nvarchar(500);not null" json:"companyName" validate:"required,min=3,max=500"`
}
