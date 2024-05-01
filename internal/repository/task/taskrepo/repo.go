package taskrepo

import (
	"context"

	"github.com/Makrorof/mongodbTutorial/internal/entity"
	"github.com/Makrorof/mongodbTutorial/internal/repository/task"
	"github.com/Makrorof/mongodbTutorial/tools"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Task struct {
	collection *mongo.Collection
}

func New(collection *mongo.Collection) task.TaskRepo {
	return &Task{
		collection: collection,
	}
}

func (r *Task) Create(ctx context.Context, task *entity.Task) (*entity.Task, error) {
	//id fix
	if task.ID.IsZero() {
		task.ID = primitive.NewObjectID()
		zap.L().Info("A new ID has been assigned because the task's ID was empty.", zap.String("new_id", task.ID.String()), zap.String("task.text[:30]", tools.StrLimit(task.Text, 30)))
	}

	zap.L().Info("Received a request to create a task.", zap.String("task.text[:30]", tools.StrLimit(task.Text, 30)))

	result, err := r.collection.InsertOne(ctx, task)

	if err != nil {
		zap.L().Error("The task creation could not be completed.", zap.Error(err), zap.String("task.text[:30]", tools.StrLimit(task.Text, 30)))
		return nil, err
	}

	zap.L().Info("The task creation has been completed.", zap.String("id", task.ID.String()), zap.String("insert_id", result.InsertedID.(primitive.ObjectID).String()), zap.String("task.text[:30]", tools.StrLimit(task.Text, 30)))

	return task, nil
}
