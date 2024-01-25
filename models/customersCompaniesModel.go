package models

import (
	"time"

	"gorm.io/gorm"
)

type CustomerCompanies struct {
	CompanyId   uint   `gorm:"primaryKey" csv:"company_id"`
	CompanyName string `csv:"company_name"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
