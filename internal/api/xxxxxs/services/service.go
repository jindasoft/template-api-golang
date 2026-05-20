package services

import (
	"context"
	"template-api-golang/internal/api/xxxxxs/models"
	"template-api-golang/internal/api/xxxxxs/repositories"

	"github.com/jindasoft/jinda-platform/xres"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type XxxxxService interface {
	ViewXxxxxByID(ctx context.Context, oid primitive.ObjectID) (*models.GetXxxxxByIDResponse, error)
	ViewXxxxxPaging(ctx context.Context, req *models.GetXxxxxPagingRequest) (*[]models.GetXxxxxPagingResponse, *xres.PagingMeta, error)
	AddXxxxx(ctx context.Context, req *models.PostXxxxxRequest) (*models.PostXxxxxResponse, error)
	EditXxxxx(ctx context.Context, oid primitive.ObjectID, req *models.PutXxxxxRequest) (*models.PutXxxxxResponse, error)
	EditXxxxxStatus(ctx context.Context, oid primitive.ObjectID, req *models.PutXxxxxStatusRequest) (*models.PutXxxxxResponse, error)
	DeleteXxxxx(ctx context.Context, oid primitive.ObjectID) error
}

type service struct {
	repo repositories.XxxxxRepository
}

func NewXxxxxService(repo repositories.XxxxxRepository) XxxxxService {
	return &service{
		repo: repo,
	}
}
