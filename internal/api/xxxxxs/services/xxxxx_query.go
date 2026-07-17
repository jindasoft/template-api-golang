package services

import (
	"context"
	"fmt"
	"template-api-golang/internal/api/xxxxxs/models"

	"github.com/jindasoft/jinda-platform/xdate"
	"github.com/jindasoft/jinda-platform/xres"
	"github.com/jindasoft/jinda-platform/xutils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *service) ViewXxxxxByID(ctx context.Context, oid primitive.ObjectID) (*models.GetXxxxxByIDResponse, error) {
	// Fetch the Xxxxx entity by ID
	entity, err := s.xxxxxRepo.FindXxxxxByID(ctx, oid)
	if err != nil {
		return nil, fmt.Errorf("failed to find Xxxxx")
	}

	// Map the entity to the response model
	res := &models.GetXxxxxByIDResponse{
		ID:          entity.ID.Hex(),
		Code:        entity.Code,
		Name:        entity.Name.LocalizeString(ctx),
		Description: entity.Description.LocalizeString(ctx),
		Icon:        entity.Icon,
		Ord:         entity.Ord,
		Status:      string(entity.Status),
		CreatedAt:   entity.CreatedAt.Format(xdate.FormatFrontend),
		UpdatedAt:   entity.UpdatedAt.Format(xdate.FormatFrontend),
	}

	return res, nil
}

func (s *service) ViewXxxxxPaging(ctx context.Context, req *models.GetXxxxxPagingRequest) (*[]models.GetXxxxxPagingResponse, *xres.PagingMeta, error) {
	offset, limit := xutils.SetDefaultPageSize(req.Offset, req.Limit)

	filter := bson.D{}
	sort := bson.D{{Key: "_id", Value: -1}}
	pt, err := s.xxxxxRepo.FindXxxxxPaging(ctx, filter, sort, offset, limit)
	if err != nil {
		return nil, nil, fmt.Errorf("error finding xxxxx list")
	}

	total, err := s.xxxxxRepo.CountXxxxx(ctx, filter)
	if err != nil {
		return nil, nil, fmt.Errorf("error counting xxxxx list: %w", err)
	}

	// Map the entities to the response model
	res := []models.GetXxxxxPagingResponse{}

	for _, p := range *pt {
		res = append(res, models.GetXxxxxPagingResponse{
			ID:   p.ID.Hex(),
			Code: p.Code,
			Name: p.Name.LocalizeString(ctx),
			Icon: p.Icon,
		})
	}

	meta := xres.PagingMeta{
		Offset: offset,
		Limit:  limit,
		Total:  total,
	}

	return &res, &meta, nil
}
