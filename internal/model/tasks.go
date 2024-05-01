package model

import "time"

type CreateTask struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Text      string    `json:"text"`
	Completed bool      `json:"completed"`
}

type UpdateTask struct {
	FilterText string `json:"filter_text"`
	Completed  bool   `json:"completed"`
	//UpdatedAt  time.Time `json:"updated_at"`
}

type Task struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Text      string    `json:"text"`
	Completed bool      `json:"completed"`
}
