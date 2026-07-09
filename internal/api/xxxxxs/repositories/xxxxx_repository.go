package repositories

import (
	"context"
	"fmt"
	"template-api-golang/internal/entities"

	"github.com/jindasoft/jinda-platform/xentities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repository) FindXxxxx(ctx context.Context, filter bson.D, sort bson.D) (*[]entities.Xxxxx, error) {
	var entity []entities.Xxxxx

	findOptions := options.Find()
	findOptions.SetSort(sort)

	if err := r.mongo.Find(ctx, filter, &entity, findOptions); err != nil {
		return nil, fmt.Errorf("error finding xxxxx list: %w", err)
	}

	return &entity, nil
}

func (r *repository) FindXxxxxByID(ctx context.Context, oid primitive.ObjectID) (*entities.Xxxxx, error) {
	var entity entities.Xxxxx
	if err := r.mongo.FindByID(ctx, oid, &entity); err != nil {
		return nil, fmt.Errorf("error finding xxxxx by ID: %w", err)
	}

	return &entity, nil
}

func (r *repository) FindXxxxxPaging(ctx context.Context, filter bson.D, sort bson.D, offset, limit int64) (*[]entities.Xxxxx, error) {
	var entity []entities.Xxxxx
	if err := r.mongo.FindPaging(ctx, filter, sort, offset, limit, &entity); err != nil {
		return nil, fmt.Errorf("error finding xxxxx list: %w", err)
	}

	return &entity, nil
}

func (r *repository) CountXxxxx(ctx context.Context, filter bson.D) (int64, error) {
	var entity entities.Xxxxx
	count, err := r.mongo.Count(ctx, filter, &entity)
	if err != nil {
		return 0, fmt.Errorf("error counting xxxxx: %w", err)
	}
	return count, nil
}

func (r *repository) InsertXxxxx(ctx context.Context, entity *entities.Xxxxx) error {
	if err := r.mongo.InsertOne(ctx, entity); err != nil {
		return fmt.Errorf("error inserting xxxxx: %w", err)
	}

	return nil
}

func (r *repository) UpdateXxxxx(ctx context.Context, entity *entities.Xxxxx) error {
	filter := bson.D{{Key: "_id", Value: entity.ID}}

	// Update the xxxxx in the database
	if err := r.mongo.UpdateOne(ctx, filter, entity); err != nil {
		return fmt.Errorf("error updating xxxxx: %w", err)
	}

	return nil
}

func (r *repository) UpdateActiveXxxxx(ctx context.Context, entity *entities.Xxxxx) error {
	filter := bson.D{
		{Key: "_id", Value: entity.ID},
		{Key: "status", Value: xentities.StatusActive},
	}

	// Update the active status of the xxxxx in the database
	if err := r.mongo.UpdateOne(ctx, filter, entity); err != nil {
		return fmt.Errorf("error updating active xxxxx: %w", err)
	}

	return nil
}

func (r *repository) UpdateInactiveXxxxx(ctx context.Context, entity *entities.Xxxxx) error {
	filter := bson.D{
		{Key: "_id", Value: entity.ID},
		{Key: "status", Value: xentities.StatusInactive},
	}

	// Update the inactive status of the xxxxx in the database
	if err := r.mongo.UpdateOne(ctx, filter, entity); err != nil {
		return fmt.Errorf("error updating inactive xxxxx: %w", err)
	}

	return nil
}

func (r *repository) SoftDeleteXxxxx(ctx context.Context, oid primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: oid}}

	var entity entities.Xxxxx
	if err := r.mongo.FindByID(ctx, oid, &entity); err != nil {
		return fmt.Errorf("error finding xxxxx by ID: %w", err)
	}

	if err := r.mongo.SoftDeleteOne(ctx, filter, &entity); err != nil {
		return fmt.Errorf("error soft deleting xxxxx: %w", err)
	}

	return nil
}
