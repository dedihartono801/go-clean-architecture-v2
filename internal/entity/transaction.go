package entity

import (
	"time"
)

type Transaction struct {
	ID               string    `json:"id"`
	AdminID          string    `json:"admin_id" validate:"required"`
	TotalQuantity    int       `json:"total_quantity" validate:"required"`
	TotalTransaction int       `json:"total_transaction" validate:"required"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
