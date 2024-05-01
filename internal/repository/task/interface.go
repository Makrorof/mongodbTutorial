package task

import (
	"context"

	"github.com/Makrorof/mongodbTutorial/internal/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskRepo interface {
	Create(ctx context.Context, task *entity.Task) (*entity.Task, error)
	GetsByFilter(ctx context.Context, filter primitive.D) ([]*entity.Task, error)
	GetAll(ctx context.Context) ([]*entity.Task, error)
	FindOneAndUpdate(ctx context.Context, filter primitive.D, update primitive.D) (*entity.Task, error)
}
