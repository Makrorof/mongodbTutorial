package itasks

import (
	"context"

	"github.com/Makrorof/mongodbTutorial/internal/model"
)

type Tasks interface {
	CreateCustom(ctx context.Context, c *model.CreateTask) (*model.Task, error)
	CreateJob(ctx context.Context, c *model.CreateTaskJob) (*model.Task, bool, error)
	Create(ctx context.Context, text string) (*model.Task, bool, error)
	GetAll(ctx context.Context) ([]*model.Task, error)
	GetPending(ctx context.Context) ([]*model.Task, error)
	GetFinished(ctx context.Context) ([]*model.Task, error)
	Update(ctx context.Context, updateTask *model.UpdateTask) error
}
