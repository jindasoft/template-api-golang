package repositories

import (
	"context"
	"template-api-examples/internal/entities"
)

func (r *repository) FindXxxxxByID(ctx context.Context, id string) (*entities.Xxxxx, error) {
	var entity entities.Xxxxx

	return &entity, nil
}
