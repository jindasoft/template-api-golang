package models

type PostXxxxxRequest struct {
	Code   string `json:"code" validate:"required"`
	NameEN string `json:"name_en" validate:"required"`
	NameTH string `json:"name_th" validate:"required"`
}

type PostXxxxxResponse struct {
	ID string `json:"id"`
}
