package models

type PostXxxxxRequest struct {
	Code        string `json:"code" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Locale      string `json:"locale" validate:"required,locale"`
}

type PostXxxxxResponse struct {
	XxxxxID string `json:"xxxxx_id"`
}
