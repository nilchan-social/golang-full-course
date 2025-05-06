package todo

import "time"

type Task struct {
	Description string
	Text        string
	IsDone      bool

	CreatedAt time.Time
	DoneAt    *time.Time
}

func NewTask(description string, text string) Task {
	return Task{
		Description: description,
		Text:        text,
		IsDone:      false,

		CreatedAt: time.Now(),
		DoneAt:    nil,
	}
}

func (t *Task) Done() {
	doneTime := time.Now()

	t.IsDone = true
	t.DoneAt = &doneTime
}
