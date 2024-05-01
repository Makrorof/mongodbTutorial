package mapper

import (
	"github.com/Makrorof/mongodbTutorial/internal/entity"
	"github.com/Makrorof/mongodbTutorial/internal/model"
)

func ToTask(t *entity.Task) *model.Task {
	return &model.Task{
		ID:        t.ID.String(),
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
		Text:      t.Text,
		Completed: t.Completed,
	}
}

func ToTasks(t []*entity.Task) []*model.Task {
	tasks := make([]*model.Task, len(t))

	for i := 0; i < len(t); i++ {
		tasks[i] = &model.Task{
			ID:        t[i].ID.String(),
			CreatedAt: t[i].CreatedAt,
			UpdatedAt: t[i].UpdatedAt,
			Text:      t[i].Text,
			Completed: t[i].Completed,
		}
	}

	return tasks
}
