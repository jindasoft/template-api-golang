package models

type GetXxxxxPagingRequest struct {
	Offset int64 `query:"offset"`
	Limit  int64 `query:"limit"`
}

type GetXxxxxPagingResponse struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}
