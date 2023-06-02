package dto

type SkuCreateDto struct {
	Name  string `json:"name" validate:"required"`
	Stock int    `json:"stock" validate:"required"`
	Price int    `json:"price" validate:"required"`
}
