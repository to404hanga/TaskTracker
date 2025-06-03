package domain

import (
	"time"

	"github.com/to404hanga/TaskTracker/internal/model"
)

type Task struct {
	Id          int
	Description string
	Status      string
	CreatedAt   string
	UpdatedAt   string
}

func FromModel(task model.Task) Task {
	return Task{
		Id:          task.Id,
		Description: task.Description,
		Status:      model.ToString(task.Status),
		CreatedAt:   time.Unix(task.CreatedAt, 0).Format(time.DateTime),
		UpdatedAt:   time.Unix(task.UpdatedAt, 0).Format(time.DateTime),
	}
}
