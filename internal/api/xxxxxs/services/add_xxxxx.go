package services

import (
	"context"
	"fmt"
	"template-api-examples/internal/api/xxxxxs/models"
	"template-api-examples/internal/entities"

	"github.com/google/uuid"
	"github.com/jindasoft/jinda-platforms/xentities"
	"github.com/jindasoft/jinda-platforms/xlogger"
	"github.com/jindasoft/jinda-platforms/xutils"
)

func (s *service) AddXxxxx(ctx context.Context, req *models.PostXxxxxRequest) (*models.PostXxxxxResponse, error) {
	method := "AddXxxxx"
	spanID := uuid.New()

	// Create a new xxxxx entity
	locale := xentities.SetLocale(req.NameEN, req.NameTH, "", "", "", "", "", "", "", "")
	entity := &entities.Xxxxx{
		Code: xutils.GetSlug(req.Code),
		Name: locale,
	}

	// Insert the new xxxxx into the database
	if err := s.repo.InsertXxxxx(ctx, entity); err != nil {
		additional := map[string]any{
			"req": xutils.JsonToStringOrDefault(req),
		}
		xlogger.AppError(ctx, method, "InsertXxxxx failed", err.Error(), &additional, &spanID)

		return nil, fmt.Errorf("failed to insert xxxxx")
	}

	// Return the response with the inserted xxxxx
	res := &models.PostXxxxxResponse{
		ID: "id",
	}

	return res, nil
}
