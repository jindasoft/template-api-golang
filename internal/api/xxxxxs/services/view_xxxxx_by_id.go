package services

import (
	"context"
	"fmt"
	"template-api-golang/internal/api/xxxxxs/models"

	"github.com/jindasoft/template-platform-go/xdate"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *service) ViewXxxxxByID(ctx context.Context, oid primitive.ObjectID) (*models.GetXxxxxByIDResponse, error) {
	// Fetch the Xxxxx entity by ID
	entity, err := s.xxxxxRepo.FindXxxxxByID(ctx, oid)
	if err != nil {
		return nil, fmt.Errorf("failed to find xxxxx by ID")
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
