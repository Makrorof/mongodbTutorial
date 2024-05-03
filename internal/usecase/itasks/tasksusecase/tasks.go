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
	"go.mongodb.org/mongo-driver/bson"
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
		ID: entity.TaskID{
			//ID:   primitive.NewObjectID(),
			Text: c.Text,
		},
		Completed: c.Completed,
	}

	zap.L().Info("A custom task creation request has been received in the use case, and a new ID has been assigned.", zap.String("new_id", et.ID.String()), zap.String("task.text[:30]", tools.StrLimit(c.Text, 30)))

	task, err := t.repo.Create(ctx, et)
	if err != nil {
		return nil, err
	}

	return mapper.ToTask(task), nil
}

// @return: task, created, err
func (t *Tasks) Create(ctx context.Context, text string) (*model.Task, bool, error) {
	et := &entity.Task{
		ID: entity.TaskID{
			//ID:   primitive.NewObjectID(),
			Text: text,
		},
		Completed: false,
	}

	zap.L().Info("A task creation request has been received in the use case, and a new ID has been assigned.", zap.String("new_id", et.ID.String()), zap.String("task.text[:30]", tools.StrLimit(text, 30)))

	task, created, err := t.repo.CreateOrUpdate(ctx, et)
	if err != nil {
		return nil, false, err
	}

	return mapper.ToTask(task), created, nil
}

func (t *Tasks) GetAll(ctx context.Context) ([]*model.Task, error) {
	tasks, err := t.repo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return mapper.ToTasks(tasks), nil
}

func (t *Tasks) Update(ctx context.Context, updateTask *model.UpdateTask) error {
	filter := bson.D{
		bson.E{
			Key: "_id",
			Value: bson.M{
				"text": updateTask.FilterText,
			},
		},
		bson.E{
			Key:   "completed",
			Value: !updateTask.Completed,
		},
	}

	update := bson.D{
		bson.E{
			Key: "$set",
			Value: bson.D{
				bson.E{
					Key:   "completed",
					Value: updateTask.Completed,
				},
				bson.E{
					Key:   "updated_at",
					Value: primitive.Timestamp{T: uint32(time.Now().Unix())},
				},
			},
		},
	}

	_, err := t.repo.FindOneAndUpdate(ctx, filter, update)

	return err
}

func (t *Tasks) GetPending(ctx context.Context) ([]*model.Task, error) {
	filter := bson.D{
		bson.E{
			Key:   "completed",
			Value: false,
		},
	}

	list, err := t.repo.GetsByFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	return mapper.ToTasks(list), nil
}
func (t *Tasks) GetFinished(ctx context.Context) ([]*model.Task, error) {
	filter := bson.D{
		bson.E{
			Key:   "completed",
			Value: true,
		},
	}

	list, err := t.repo.GetsByFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	return mapper.ToTasks(list), nil
}
