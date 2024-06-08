package web

type MerchantRequest struct {
	Name string `json:"name" validate:"required"`
}
