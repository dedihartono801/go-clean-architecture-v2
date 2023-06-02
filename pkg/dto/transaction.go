package dto

type TransactionCheckoutDto struct {
	Items []TransactionItemDto `json:"items" validate:"required"`
}

type TransactionItemDto struct {
	ID       string `json:"id" `
	Quantity int    `json:"quantity" `
}
