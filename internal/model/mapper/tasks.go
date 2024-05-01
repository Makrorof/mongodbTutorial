package mapper

import (
	"github.com/Makrorof/mongodbTutorial/internal/entity"
	"github.com/Makrorof/mongodbTutorial/internal/model"
)

func ToTask(t *entity.Task) *model.Task {
	return &model.Task{
		ID: model.TaskID{
			//ID:   t.ID.ID.String(),
			Text: t.ID.Text,
		},
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
		Completed: t.Completed,
	}
}

func ToTasks(t []*entity.Task) []*model.Task {
	tasks := make([]*model.Task, len(t))

	for i := 0; i < len(t); i++ {
		tasks[i] = &model.Task{
			ID: model.TaskID{
				//ID:   t[i].ID.ID.String(),
				Text: t[i].ID.Text,
			},
			CreatedAt: t[i].CreatedAt,
			UpdatedAt: t[i].UpdatedAt,
			Completed: t[i].Completed,
		}
	}

	return tasks
}
