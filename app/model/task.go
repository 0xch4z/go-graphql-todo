package model

import "time"

// Task type
type Task struct {
	ID          uint       `json:"id" gorm:"primary_key"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt"`
	UserID      uint       `json:"userId"`
	IsComplete  bool       `json:"isComplete"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
}

// Complete completes a given task
func (t *Task) Complete() {
	t.IsComplete = true
}

// Undo undoes a given task
func (t *Task) Undo() {
	t.IsComplete = false
}
