package models

import "github.com/jindasoft/jinda-platform/xres"

type SwaggerGetXxxxxByIDResponse struct {
	xres.SuccessResponse

	Data GetXxxxxByIDResponse `json:"data"`
}

type GetXxxxxByIDResponse struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}
