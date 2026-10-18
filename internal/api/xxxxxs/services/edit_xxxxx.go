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

const (
	newMongoLocaleFailedMsg    = "NewMongoLocale failed"
	failedToCreateLocaleErrMsg = "failed to create locale"
	findXxxxxByIDFailedMsg     = "FindXxxxxByID failed"
	failedToFindXxxxxErrMsg    = "failed to find Xxxxx"
	insertXxxxxFailedMsg       = "InsertXxxxx failed"
	updateXxxxxFailedMsg       = "UpdateXxxxx failed"
	softDeleteXxxxxFailedMsg   = "SoftDeleteXxxxx failed"
	failedToSoftDeleteXxxxxMsg = "failed to soft delete Xxxxx"
	failedToInsertXxxxxMsg     = "failed to insert Xxxxx"
	failedToUpdateXxxxxMsg     = "failed to update Xxxxx"
)

func (s *service) EditXxxxx(ctx context.Context, oid primitive.ObjectID, req *models.PutXxxxxRequest) (*models.PutXxxxxResponse, error) {
	method := "EditXxxxx"
	spanID := uuid.New()

	// Check if the Xxxxx exists
	xxxxx, err := s.xxxxxRepo.FindXxxxxByID(ctx, oid)
	if err != nil {
		additional := map[string]any{
			"req": xutils.JsonToStringOrDefault(req),
		}
		xlogger.AppError(ctx, method, findXxxxxByIDFailedMsg, err.Error(), &additional, &spanID)

		return nil, errors.New(failedToFindXxxxxErrMsg)
	}

	nameLocale, err := xentities.NewMongoLocale(req.Locale, req.Name)
	if err != nil {
		additional := map[string]any{
			"req": xutils.JsonToStringOrDefault(req),
		}
		xlogger.AppError(ctx, method, newMongoLocaleFailedMsg, err.Error(), &additional, &spanID)

		return nil, errors.New(failedToCreateLocaleErrMsg)
	}

	descLocale, err := xentities.NewMongoLocale(req.Locale, req.Description)
	if err != nil {
		additional := map[string]any{
			"req": xutils.JsonToStringOrDefault(req),
		}
		xlogger.AppError(ctx, method, newMongoLocaleFailedMsg, err.Error(), &additional, &spanID)

		return nil, errors.New(failedToCreateLocaleErrMsg)
	}

	// Update the Xxxxx with the new values
	xxxxx.Code = xutils.GetSlug(req.Code)
	xxxxx.Name = nameLocale
	xxxxx.Description = descLocale
	xxxxx.Ord = req.Ord

	// Save the updated Xxxxx back to the database
	if err := s.xxxxxRepo.UpdateXxxxx(ctx, xxxxx); err != nil {
		additional := map[string]any{
			"req": xutils.JsonToStringOrDefault(req),
		}
		spanID := uuid.New()
		xlogger.AppError(ctx, method, updateXxxxxFailedMsg, err.Error(), &additional, &spanID)

		return nil, errors.New(failedToUpdateXxxxxMsg)
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
