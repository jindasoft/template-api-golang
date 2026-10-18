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
		subject := "failed to find xxxxx by ID"
		additional := map[string]any{
			"oid": xutils.JsonToStringOrDefault(oid),
		}
		xlogger.AppWarn(ctx, method, subject, err.Error(), &additional, &spanID)

		return errors.New(subject)
	}

	// Soft delete the Xxxxx
	if err := s.xxxxxRepo.SoftDeleteXxxxx(ctx, pt.ID); err != nil {
		subject := "failed to soft delete xxxxx"
		additional := map[string]any{
			"oid": xutils.JsonToStringOrDefault(oid),
		}
		xlogger.AppError(ctx, method, subject, err.Error(), &additional, &spanID)

		return errors.New(subject)
	}

	return nil
}
