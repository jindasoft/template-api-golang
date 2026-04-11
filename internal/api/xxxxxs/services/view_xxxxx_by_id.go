package services

import (
	"context"
	"fmt"
	"template-api-examples/internal/api/xxxxxs/models"

	"github.com/jindasoft/jinda-platforms/xdate"
)

func (s *service) ViewXxxxxByID(ctx context.Context, id string) (*models.GetXxxxxByIDResponse, error) {
	// Fetch the xxxxx entity by ID
	entity, err := s.repo.FindXxxxxByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find xxxxx")
	}

	// Return the response
	res := &models.GetXxxxxByIDResponse{
		ID:        "id", // entity.ID --- IGNORE ---
		Code:      entity.Code,
		Name:      "entity.Name",
		Status:    "entity.Status",
		CreatedAt: entity.CreatedAt.Format(xdate.FormatISO8601),
		UpdatedAt: entity.UpdatedAt.Format(xdate.FormatISO8601),
	}

	return res, nil
}
