package repositories

import (
	"context"
	"template-api-examples/internal/entities"
)

type XxxxxRepository interface {
	FindXxxxxByID(ctx context.Context, id string) (*entities.Xxxxx, error)
	FindXxxxxPaging(ctx context.Context) (*[]entities.Xxxxx, error)
	InsertXxxxx(ctx context.Context, entity *entities.Xxxxx) error
}

type repository struct {
}

func NewXxxxxRepository() XxxxxRepository {
	return &repository{}
}
