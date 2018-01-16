package model

// Task type
type Task struct {
	Model
	UserID      uint   `json:"userId"`
	IsComplete  bool   `json:"isComplete"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Complete completes a given task
func (t *Task) Complete() {
	t.IsComplete = true
}

// Undo undoes a given task
func (t *Task) Undo() {
	t.IsComplete = false
}
