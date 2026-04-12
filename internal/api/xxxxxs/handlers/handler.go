package handlers

import "template-api-golang/internal/api/xxxxxs/services"

type handler struct {
	service services.XxxxxService
}

func NewHandler(service services.XxxxxService) *handler {
	return &handler{
		service: service,
	}
}
