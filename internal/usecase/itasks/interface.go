package itasks

import (
	"context"

	"github.com/Makrorof/mongodbTutorial/internal/model"
)

type Tasks interface {
	CreateCustom(ctx context.Context, c *model.CreateTask) (*model.Task, error)
	Create(ctx context.Context, text string) (*model.Task, error)
	GetAll(ctx context.Context) ([]*model.Task, error)
}
