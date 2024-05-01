package tasksusecase

import (
	"context"
	"time"

	"github.com/Makrorof/mongodbTutorial/internal/entity"
	"github.com/Makrorof/mongodbTutorial/internal/model"
	"github.com/Makrorof/mongodbTutorial/internal/model/mapper"
	"github.com/Makrorof/mongodbTutorial/internal/repository/task"
	itask "github.com/Makrorof/mongodbTutorial/internal/usecase/itasks"
	"github.com/Makrorof/mongodbTutorial/tools"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type Tasks struct {
	repo task.TaskRepo
}

func New(repo task.TaskRepo) itask.Tasks {
	return &Tasks{
		repo: repo,
	}
}

func (t *Tasks) CreateCustom(ctx context.Context, c *model.CreateTask) (*model.Task, error) {
	et := &entity.Task{
		ID:        primitive.NewObjectID(),
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		Text:      c.Text,
		Completed: c.Completed,
	}

	zap.L().Info("A custom task creation request has been received in the use case, and a new ID has been assigned.", zap.String("new_id", et.ID.String()), zap.String("task.text[:30]", tools.StrLimit(c.Text, 30)))

	task, err := t.repo.Create(ctx, et)
	if err != nil {
		return nil, err
	}

	return mapper.ToTask(task), nil
}

func (t *Tasks) Create(ctx context.Context, text string) (*model.Task, error) {
	et := &entity.Task{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Text:      text,
		Completed: false,
	}

	zap.L().Info("A task creation request has been received in the use case, and a new ID has been assigned.", zap.String("new_id", et.ID.String()), zap.String("task.text[:30]", tools.StrLimit(text, 30)))

	task, err := t.repo.Create(ctx, et)
	if err != nil {
		return nil, err
	}

	return mapper.ToTask(task), nil
}
