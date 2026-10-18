package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jindasoft/template-platform-go/xlogger"
	"github.com/jindasoft/template-platform-go/xutils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *service) DeleteXxxxx(ctx context.Context, oid primitive.ObjectID) error {
	method := "DeleteXxxxx"
	spanID := uuid.New()

	// Check if the Xxxxx exists
	pt, err := s.xxxxxRepo.FindXxxxxByID(ctx, oid)
	if err != nil {
		additional := map[string]any{
			"oid": xutils.JsonToStringOrDefault(oid),
		}
		xlogger.AppWarn(ctx, method, findXxxxxByIDFailedMsg, err.Error(), &additional, &spanID)

		return errors.New(failedToFindXxxxxErrMsg)
	}

	// Soft delete the Xxxxx
	if err := s.xxxxxRepo.SoftDeleteXxxxx(ctx, pt.ID); err != nil {
		additional := map[string]any{
			"oid": xutils.JsonToStringOrDefault(oid),
		}
		xlogger.AppError(ctx, method, softDeleteXxxxxFailedMsg, err.Error(), &additional, &spanID)

		return errors.New(failedToSoftDeleteXxxxxMsg)
	}

	return nil
}
