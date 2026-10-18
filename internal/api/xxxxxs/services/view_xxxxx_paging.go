package services

import (
	"context"
	"fmt"
	"template-api-golang/internal/api/xxxxxs/models"

	"github.com/jindasoft/template-platform-go/xres"
	"github.com/jindasoft/template-platform-go/xutils"
	"go.mongodb.org/mongo-driver/bson"
)

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
