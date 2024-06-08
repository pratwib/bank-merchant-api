package web

type AddBalanceRequest struct {
	CustomerID string  `json:"customer_id" validate:"required"`
	Amount     float64 `json:"amount" validate:"required,gt=0"`
}
