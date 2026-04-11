package models

import "github.com/jindasoft/jinda-platforms/xres"

type GetXxxxxPagingRequest struct {
	Limit  int64 `query:"limit"`
	Offset int64 `query:"offset"`
}

// For Swagger documentation (expanded structure)
type SwaggerGetXxxxxPagingResponse struct {
	xres.PagingResponse

	Data []GetXxxxxPagingResponse `json:"data"`
}

type GetXxxxxPagingResponse struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
