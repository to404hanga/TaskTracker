package model

type Task struct {
	Id          int        `json:"id"`          // A unique identifier for the task.
	Description string     `json:"description"` // A short description of the task.
	Status      StatusCode `json:"status"`      // The status of the task. (todo, in-progress, done)
	CreatedAt   int64      `json:"createdAt"`   // The date and time when the task was created.
	UpdatedAt   int64      `json:"updatedAt"`   // The date and time when the task was last updated.
}
