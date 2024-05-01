package task

import (
	"context"

	"github.com/Makrorof/mongodbTutorial/internal/entity"
)

type TaskRepo interface {
	Create(ctx context.Context, task *entity.Task) (*entity.Task, error)
}
