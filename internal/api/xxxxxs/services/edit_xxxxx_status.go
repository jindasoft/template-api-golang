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

func (s *service) EditXxxxxStatus(ctx context.Context, oid primitive.ObjectID, req *models.PutXxxxxStatusRequest) (*models.PutXxxxxResponse, error) {
	method := "EditXxxxxStatus"
	spanID := uuid.New()

	// Check if the Xxxxx exists
	xxxxx, err := s.xxxxxRepo.FindXxxxxByID(ctx, oid)
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
	if err := s.xxxxxRepo.UpdateXxxxx(ctx, xxxxx); err != nil {
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
