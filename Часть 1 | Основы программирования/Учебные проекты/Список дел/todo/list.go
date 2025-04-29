package todo

type List struct {
	tasks map[string]Task
}

func NewList() *List {
	return &List{
		tasks: make(map[string]Task),
	}
}

func (l *List) AddTask(task Task) {
	l.tasks[task.Description] = task
}

func (l *List) ListTasks() map[string]Task {
	return l.tasks
}

func (l *List) DoneTask(description string) string {
	task, ok := l.tasks[description]
	if !ok {
		return taskNotFound
	}

	task.Done()

	l.tasks[description] = task

	return ""
}

func (l *List) DeleteTask(description string) string {
	_, ok := l.tasks[description]
	if !ok {
		return taskNotFound
	}

	delete(l.tasks, description)

	return ""
}
