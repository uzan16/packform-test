package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	ID         uint      `gorm:"primaryKey" csv:"id"`
	OrderName  string    `csv:"order_name"`
	CustomerID string    `csv:"customer_id"`
	CreatedAt  time.Time `csv:"created_at"`
}

func BuildGetOrderQuery(db *gorm.DB) *gorm.DB {
	subQuery := db.Model(&OrderItems{}).
		Select("order_items.*, coalesce (SUM(deliveries.delivered_quantity), 0) as delivered_quantity").
		Joins("LEFT JOIN deliveries ON deliveries.order_item_id = order_items.id").
		Group("order_items.id")

	query := db.Model(&Orders{}).
		Select(fmt.Sprint(
			"orders.id,",
			"orders.order_name,",
			"customer_companies.company_name,",
			"customers.name AS customer_name,",
			"orders.created_at,",
			"SUM(order_items.price_per_unit * order_items.quantity) as total_amount,",
			"SUM(order_items.price_per_unit * order_items.delivered_quantity) as delivered_amount,",
			"string_agg(order_items.product, ', ' ORDER BY order_items.product) AS products",
		)).
		Joins("INNER JOIN customers ON customers.user_id = orders.customer_id").
		Joins("INNER JOIN customer_companies ON customer_companies.company_id = customers.company_id").
		Joins("LEFT JOIN (?) as order_items ON order_items.order_id = orders.id", subQuery).
		Group("orders.id, customer_companies.company_id, customers.user_id")
	return query
}
