package models

type GetXxxxxPagingRequest struct {
	Limit  int64 `query:"limit"`
	Offset int64 `query:"offset"`
}

type GetXxxxxPagingResponse struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}
