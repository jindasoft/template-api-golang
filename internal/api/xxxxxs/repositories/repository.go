package repositories

import (
	"context"
	"template-api-golang/internal/entities"

	"github.com/jindasoft/jinda-platform/xdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type XxxxxRepository interface {
	FindXxxxxByID(ctx context.Context, oid primitive.ObjectID) (*entities.Xxxxx, error)
	FindXxxxxPaging(ctx context.Context, filter bson.D, sort bson.D, offset, limit int64) (*[]entities.Xxxxx, error)
	FindXxxxx(ctx context.Context, filter bson.D, sort bson.D) (*[]entities.Xxxxx, error)
	CountXxxxx(ctx context.Context, filter bson.D) (int64, error)
	InsertXxxxx(ctx context.Context, entity *entities.Xxxxx) error
	UpdateXxxxx(ctx context.Context, entity *entities.Xxxxx) error
	UpdateActiveXxxxx(ctx context.Context, entity *entities.Xxxxx) error
	UpdateInactiveXxxxx(ctx context.Context, entity *entities.Xxxxx) error
	SoftDeleteXxxxx(ctx context.Context, oid primitive.ObjectID) error
}

type repository struct {
	mongo xdb.MongoService
}

func NewXxxxxRepository(mongo xdb.MongoService) XxxxxRepository {
	return &repository{
		mongo: mongo,
	}
}
