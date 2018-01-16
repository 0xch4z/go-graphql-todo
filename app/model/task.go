package model

// Task type
type Task struct {
	Model
	OwnerID     uint   `json:"ownerId" sql:"name:ownerId"`
	IsComplete  bool   `json:"isComplete" sql:"name:isComplete"`
	Title       string `json:"title" sql:"name:title"`
	Description string `json:"description" sql:"name:description"`
}

// Complete completes a given task
func (t *Task) Complete() {
	t.IsComplete = true
}

// Undo undoes a given task
func (t *Task) Undo() {
	t.IsComplete = false
}
