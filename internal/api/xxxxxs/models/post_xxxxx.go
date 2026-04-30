package models

import "github.com/jindasoft/jinda-platform/xres"

type PostXxxxxRequest struct {
	Code   string `json:"code" validate:"required"`
	NameEN string `json:"name_en" validate:"required"`
	NameTH string `json:"name_th" validate:"required"`
}

type SwaggerPostXxxxxResponse struct {
	xres.CreatedResponse

	Data PostXxxxxResponse `json:"data"`
}
type PostXxxxxResponse struct {
	ID string `json:"id"`
}
