package entity

import (
	"time"
)

type Sku struct {
	ID        string    `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Stock     int       `json:"stock" validate:"required"`
	Price     int       `json:"price" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
