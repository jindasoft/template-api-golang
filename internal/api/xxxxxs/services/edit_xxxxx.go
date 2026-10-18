package services

import (
	"context"
	"errors"
	"template-api-golang/internal/api/xxxxxs/models"

	"github.com/google/uuid"
	"github.com/jindasoft/template-platform-go/xentities"
	"github.com/jindasoft/template-platform-go/xlogger"
	"github.com/jindasoft/template-platform-go/xutils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *service) EditXxxxx(ctx context.Context, oid primitive.ObjectID, req *models.PutXxxxxRequest) (*models.PutXxxxxResponse, error) {
	method := "EditXxxxx"
	spanID := uuid.New()

	// Check if the Xxxxx exists
	xxxxx, err := s.xxxxxRepo.FindXxxxxByID(ctx, oid)
	if err != nil {
		subject := "failed to find xxxxx by ID"
		additional := map[string]any{
			"req": xutils.JsonToStringOrDefault(req),
		}
		xlogger.AppError(ctx, method, subject, err.Error(), &additional, &spanID)

		return nil, errors.New(subject)
	}

	nameLocale, err := xentities.NewMongoLocale(req.Locale, req.Name)
	if err != nil {
		subject := "failed to create name locale"
		additional := map[string]any{
			"req": xutils.JsonToStringOrDefault(req),
		}
		xlogger.AppError(ctx, method, subject, err.Error(), &additional, &spanID)

		return nil, errors.New(subject)
	}

	descLocale, err := xentities.NewMongoLocale(req.Locale, req.Description)
	if err != nil {
		subject := "failed to create description locale"
		additional := map[string]any{
			"req": xutils.JsonToStringOrDefault(req),
		}
		xlogger.AppError(ctx, method, subject, err.Error(), &additional, &spanID)

		return nil, errors.New(subject)
	}

	// Update the Xxxxx with the new values
	xxxxx.Code = xutils.GetSlug(req.Code)
	xxxxx.Name = nameLocale
	xxxxx.Description = descLocale
	xxxxx.Ord = req.Ord

	// Save the updated Xxxxx back to the database
	if err := s.xxxxxRepo.UpdateXxxxx(ctx, xxxxx); err != nil {
		subject := "failed to update xxxxx"
		additional := map[string]any{
			"req": xutils.JsonToStringOrDefault(req),
		}
		spanID := uuid.New()
		xlogger.AppError(ctx, method, subject, err.Error(), &additional, &spanID)

		return nil, errors.New(subject)
	}

	// Return the response with the updated Xxxxx
	res := &models.PutXxxxxResponse{
		ID:          xxxxx.ID.Hex(),
		Code:        xxxxx.Code,
		Name:        xxxxx.Name.LocalizeString(ctx),
		Description: xxxxx.Description.LocalizeString(ctx),
		Ord:         xxxxx.Ord,
		Status:      string(xxxxx.Status),
	}

	return res, nil
}
