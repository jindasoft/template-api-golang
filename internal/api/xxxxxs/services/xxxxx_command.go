package services

import (
	"context"
	"errors"
	"template-api-golang/internal/api/xxxxxs/models"
	"template-api-golang/internal/entities"

	"github.com/google/uuid"
	"github.com/jindasoft/jinda-platform/xentities"
	"github.com/jindasoft/jinda-platform/xlogger"
	"github.com/jindasoft/jinda-platform/xutils"
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

func (s *service) AddXxxxx(ctx context.Context, req *models.PostXxxxxRequest) (*models.PostXxxxxResponse, error) {
	method := "AddXxxxx"
	spanID := uuid.New()

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

	// Create a new Xxxxx entity
	entity := &entities.Xxxxx{
		Code:        xutils.GetSlug(req.Code),
		Name:        nameLocale,
		Description: descLocale,
	}

	// Insert the new Xxxxx into the database
	if err := s.repo.InsertXxxxx(ctx, entity); err != nil {
		additional := map[string]any{
			"req": xutils.JsonToStringOrDefault(req),
		}
		spanID := uuid.New()
		xlogger.AppError(ctx, method, insertXxxxxFailedMsg, err.Error(), &additional, &spanID)

		return nil, errors.New(failedToInsertXxxxxMsg)
	}

	// Return the response with the inserted Xxxxx
	res := &models.PostXxxxxResponse{
		XxxxxID: entity.ID.Hex(),
	}

	return res, nil
}

func (s *service) EditXxxxx(ctx context.Context, oid primitive.ObjectID, req *models.PutXxxxxRequest) (*models.PutXxxxxResponse, error) {
	method := "EditXxxxx"
	spanID := uuid.New()

	// Check if the Xxxxx exists
	xxxxx, err := s.repo.FindXxxxxByID(ctx, oid)
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
	if err := s.repo.UpdateXxxxx(ctx, xxxxx); err != nil {
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

func (s *service) EditXxxxxStatus(ctx context.Context, oid primitive.ObjectID, req *models.PutXxxxxStatusRequest) (*models.PutXxxxxResponse, error) {
	method := "EditXxxxxStatus"
	spanID := uuid.New()

	// Check if the Xxxxx exists
	xxxxx, err := s.repo.FindXxxxxByID(ctx, oid)
	if err != nil {
		additional := map[string]any{
			"oid": xutils.JsonToStringOrDefault(oid),
		}
		xlogger.AppWarn(ctx, method, findXxxxxByIDFailedMsg, err.Error(), &additional, &spanID)

		return nil, errors.New(failedToFindXxxxxErrMsg)
	}

	// Update the active status of the Xxxxx
	status, ok := xentities.FromStatusCode(req.Status)
	if !ok {
		return nil, errors.New("invalid status")
	}
	xxxxx.Status = status

	// Save the updated Xxxxx back to the database
	if err := s.repo.UpdateXxxxx(ctx, xxxxx); err != nil {
		additional := map[string]any{
			"oid": xutils.JsonToStringOrDefault(oid),
		}
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

func (s *service) DeleteXxxxx(ctx context.Context, oid primitive.ObjectID) error {
	method := "DeleteXxxxx"
	spanID := uuid.New()

	// Check if the Xxxxx exists
	pt, err := s.repo.FindXxxxxByID(ctx, oid)
	if err != nil {
		additional := map[string]any{
			"oid": xutils.JsonToStringOrDefault(oid),
		}
		xlogger.AppWarn(ctx, method, findXxxxxByIDFailedMsg, err.Error(), &additional, &spanID)

		return errors.New(failedToFindXxxxxErrMsg)
	}

	// Soft delete the Xxxxx
	if err := s.repo.SoftDeleteXxxxx(ctx, pt.ID); err != nil {
		additional := map[string]any{
			"oid": xutils.JsonToStringOrDefault(oid),
		}
		xlogger.AppError(ctx, method, softDeleteXxxxxFailedMsg, err.Error(), &additional, &spanID)

		return errors.New(failedToSoftDeleteXxxxxMsg)
	}

	return nil
}
