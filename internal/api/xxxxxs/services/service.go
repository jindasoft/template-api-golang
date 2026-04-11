package services

import (
	"context"
	"template-api-examples/internal/api/xxxxxs/models"
	"template-api-examples/internal/api/xxxxxs/repositories"
)

type XxxxxService interface {
	ViewXxxxxPaging(ctx context.Context, req *models.GetXxxxxPagingRequest) (*[]models.GetXxxxxPagingResponse, error)
	ViewXxxxxByID(ctx context.Context, id string) (*models.GetXxxxxByIDResponse, error)
	AddXxxxx(ctx context.Context, req *models.PostXxxxxRequest) (*models.PostXxxxxResponse, error)
}

type service struct {
	repo repositories.XxxxxRepository
}

func NewXxxxxService(repo repositories.XxxxxRepository) XxxxxService {
	return &service{
		repo: repo,
	}
}
