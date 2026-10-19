package services

import (
	"context"
	"errors"
	"template-api-golang/internal/api/xxxxxs/models"
	"template-api-golang/internal/entities"

	"github.com/google/uuid"
	"github.com/jindasoft/template-platform-go/xentities"
	"github.com/jindasoft/template-platform-go/xlogger"
	"github.com/jindasoft/template-platform-go/xutils"
)

func (s *service) AddXxxxx(ctx context.Context, req *models.PostXxxxxRequest) (*models.PostXxxxxResponse, error) {
	method := "AddXxxxx"
	spanID := uuid.New()

	nameLocale, err := xentities.NewMongoLocale(req.Locale, req.Name)
	if err != nil {
		subject := "failed to create locale"
		additional := map[string]any{
			"req": xutils.JsonToStringOrDefault(req),
		}
		xlogger.AppError(ctx, method, subject, err.Error(), &additional, &spanID)

		return nil, errors.New(subject)
	}

	descLocale, err := xentities.NewMongoLocale(req.Locale, req.Description)
	if err != nil {
		subject := "failed to create locale"
		additional := map[string]any{
			"req": xutils.JsonToStringOrDefault(req),
		}
		xlogger.AppError(ctx, method, subject, err.Error(), &additional, &spanID)

		return nil, errors.New(subject)
	}

	// Create a new Xxxxx entity
	entity := &entities.Xxxxx{
		Code:        xutils.GetSlug(req.Code),
		Name:        nameLocale,
		Description: descLocale,
	}

	// Insert the new Xxxxx into the database
	if err := s.xxxxxRepo.InsertXxxxx(ctx, entity); err != nil {
		subject := "failed to insert xxxxx"
		additional := map[string]any{
			"req": xutils.JsonToStringOrDefault(req),
		}
		spanID := uuid.New()
		xlogger.AppError(ctx, method, subject, err.Error(), &additional, &spanID)

		return nil, errors.New(subject)
	}

	// Return the response with the inserted Xxxxx
	res := &models.PostXxxxxResponse{
		XxxxxID: entity.ID.Hex(),
	}

	return res, nil
}
