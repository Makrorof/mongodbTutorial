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
