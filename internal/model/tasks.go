package model

import (
	"time"
)

type CreateTask struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Text      string    `json:"text"`
	Completed bool      `json:"completed"`
	Jobs      []Job     `json:"jobs"`
}

type Job struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

type CreateTaskJob struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Text      string    `json:"text"`
	Completed bool      `json:"completed"`
	Jobs      []Job     `json:"jobs"`
}

type UpdateTask struct {
	FilterText string `json:"filter_text"`
	Completed  bool   `json:"completed"`
	//UpdatedAt  time.Time `json:"updated_at"`
}

type Task struct {
	ID        TaskID    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Text      string    `json:"text"`
	Completed bool      `json:"completed"`
}

type TaskID struct {
	//ID   string `bson:"id" json:"id"`
	Text string `bson:"text" json:"text"`
}
