package models

type PutXxxxxRequest struct {
	Code        string `json:"code" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Ord         int16  `json:"ord" validate:"required"`
	Locale      string `json:"locale" validate:"required,locale"`
}

type PutXxxxxStatusRequest struct {
	Status string `json:"status" validate:"required"`
}

type PutXxxxxResponse struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Ord         int16  `json:"ord"`
	Status      string `json:"status"`
}
