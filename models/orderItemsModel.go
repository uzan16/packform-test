package models

import "gorm.io/gorm"

type OrderItems struct {
	gorm.Model
	ID           uint    `gorm:"primaryKey" csv:"id"`
	OrderID      uint    `csv:"order_id"`
	PricePerUnit float32 `csv:"price_per_unit"`
	Quantity     uint    `csv:"quantity"`
	Product      string  `csv:"product"`
}
