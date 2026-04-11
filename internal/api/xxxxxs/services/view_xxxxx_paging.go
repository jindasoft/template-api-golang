package services

import (
	"context"
	"fmt"
	"template-api-examples/internal/api/xxxxxs/models"

	"github.com/jindasoft/jinda-platforms/xdate"
)

func (s *service) ViewXxxxxPaging(ctx context.Context, req *models.GetXxxxxPagingRequest) (*[]models.GetXxxxxPagingResponse, error) {

	pt, err := s.repo.FindXxxxxPaging(ctx)
	if err != nil {
		return nil, fmt.Errorf("error finding xxxxx list")
	}

	// Return the response
	res := []models.GetXxxxxPagingResponse{}

	for _, p := range *pt {
		res = append(res, models.GetXxxxxPagingResponse{
			ID:        "id", // p.ID --- IGNORE ---
			Code:      p.Code,
			Name:      p.Name.String(ctx),
			Status:    "p.Status",
			CreatedAt: p.CreatedAt.Format(xdate.FormatISO8601),
			UpdatedAt: p.UpdatedAt.Format(xdate.FormatISO8601),
		})
	}

	return &res, nil
}
