package dto

import "time"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PaginatedResponse struct {
	Response
	TotalRows int64 `json:"totalRows"`
}

type OrderResponse struct {
	ID              uint      `json:"id"`
	OrderName       string    `json:"orderName"`
	CreatedAt       time.Time `json:"createdAt"`
	CompanyName     string    `json:"companyName"`
	CustomerName    string    `json:"customerName"`
	TotalAmount     float32   `json:"totalAmount"`
	DeliveredAmount float32   `json:"deliveredAmount"`
	Products        string    `json:"products"`
}
