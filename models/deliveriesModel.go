package models

import "gorm.io/gorm"

type Deliveries struct {
	gorm.Model
	ID                uint `gorm:"primaryKey" csv:"id"`
	OrderItemID       uint `csv:"order_item_id"`
	DeliveredQuantity uint `csv:"delivered_quantity"`
}
