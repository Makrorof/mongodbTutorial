package taskrepo

import (
	"context"
	"time"

	"github.com/Makrorof/mongodbTutorial/internal/entity"
	"github.com/Makrorof/mongodbTutorial/internal/repository/task"
	"github.com/Makrorof/mongodbTutorial/tools"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	//if task.ID.ID.IsZero() {
	//	task.ID.ID = primitive.NewObjectID()
	//	zap.L().Info("A new ID has been assigned because the task's ID was empty.", zap.String("new_id", task.ID.ID.String()), zap.String("task.text[:30]", tools.StrLimit(task.ID.Text, 30)))
	//}

	task.CreatedAt = primitive.Timestamp{T: uint32(time.Now().Unix())}
	task.UpdatedAt = primitive.Timestamp{T: uint32(time.Now().Unix())}

	zap.L().Info("Received a request to create a task.", zap.String("task.text[:30]", tools.StrLimit(task.ID.Text, 30)))

	result, err := r.collection.InsertOne(ctx, task)

	if err != nil {
		zap.L().Error("The task creation could not be completed.", zap.Error(err), zap.String("task.text[:30]", tools.StrLimit(task.ID.Text, 30)))
		return nil, err
	}

	zap.L().Info("The task creation has been completed.", zap.String("id", task.ID.String()), zap.String("insert_id", entity.ParseTaskID(result.InsertedID).String()), zap.String("task.text[:30]", tools.StrLimit(task.ID.Text, 30)))

	return task, nil
}

func (r *Task) CreateOrUpdate(ctx context.Context, task *entity.Task) (*entity.Task, bool, error) {
	//id fix
	//if task.ID.ID.IsZero() {
	//	task.ID.ID = primitive.NewObjectID()
	//	zap.L().Info("A new ID has been assigned because the task's ID was empty.", zap.String("new_id", task.ID.ID.String()), zap.String("task.text[:30]", tools.StrLimit(task.ID.Text, 30)))
	//}

	task.CreatedAt = primitive.Timestamp{T: uint32(time.Now().Unix())}
	task.UpdatedAt = primitive.Timestamp{T: uint32(time.Now().Unix())}

	zap.L().Info("Received a request to create a task.", zap.String("task.text[:30]", tools.StrLimit(task.ID.Text, 30)))

	filter := bson.M{
		"_id": task.ID,
	}

	update := bson.D{
		bson.E{
			Key: "$set",
			Value: bson.D{
				bson.E{
					Key:   "updated_at",
					Value: task.UpdatedAt,
				},
				bson.E{
					Key:   "completed",
					Value: task.Completed,
				},
			},
		},
		bson.E{
			Key: "$setOnInsert",
			Value: bson.D{
				bson.E{
					Key:   "created_at",
					Value: task.CreatedAt,
				},
			},
		},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))

	//result, err := r.collection.InsertOne(ctx, task)

	if err != nil {
		zap.L().Error("The task creation could not be completed.", zap.Error(err), zap.String("task.text[:30]", tools.StrLimit(task.ID.Text, 30)))
		return nil, false, err
	}

	zap.L().Info("The task creation has been completed.", zap.String("id", task.ID.String()), zap.String("insert_id", entity.ParseTaskID(result.UpsertedID).String()), zap.String("task.text[:30]", tools.StrLimit(task.ID.Text, 30)))

	return task, result.UpsertedCount > 0, nil
}

func (r *Task) GetsByFilter(ctx context.Context, filter primitive.D) ([]*entity.Task, error) {
	tasks := make([]*entity.Task, 0)

	zap.L().Info("A request to retrieve tasks based on a filter has been received.")

	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		zap.L().Error("There was an issue while fetching the tasks. [get cursor]", zap.Error(err))
		return tasks, err
	}

	for cur.Next(ctx) {
		var t entity.Task // =new //hata verirse gereksiz olusturma olur
		err := cur.Decode(&t)
		if err != nil {
			zap.L().Error("There was an issue while fetching the tasks.", zap.Error(err))
			return tasks, err
		}

		tasks = append(tasks, &t)
	}

	zap.L().Info("The tasks were successfully fetched.", zap.Int("len", len(tasks)))

	return tasks, nil
}

func (r *Task) GetAll(ctx context.Context) ([]*entity.Task, error) {
	return r.GetsByFilter(ctx, bson.D{{}})
}

func (r *Task) FindOneAndUpdate(ctx context.Context, filter primitive.D, update primitive.D) (*entity.Task, error) {
	t := &entity.Task{}
	if err := r.collection.FindOneAndUpdate(ctx, filter, update).Decode(t); err != nil {
		return nil, err
	}

	return t, nil
}
