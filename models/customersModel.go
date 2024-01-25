package models

import (
	"time"

	"gorm.io/gorm"
)

type Customers struct {
	UserId      string `gorm:"primaryKey" csv:"user_id"`
	Login       string `csv:"login"`
	Password    string `csv:"password"`
	Name        string `csv:"name"`
	CompanyId   uint   `csv:"company_id"`
	CreditCards string `csv:"credit_cards"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
