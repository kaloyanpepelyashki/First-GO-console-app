package taskmanager

type Task struct {
	id          int
	title       string
	description string
	completed   bool
	deleted     bool
	endDate     string
}

func (t *Task) CreateTask(id int, title, description string, endDate string) {
	t.id = id
	t.title = title
	t.description = description
	t.completed = false
	t.deleted = false
	t.endDate = endDate
}

func (t *Task) UpdateTask(title, description string) {
	t.title = title
	t.description = description
}

func (t *Task) DeleteTask() {
	t.id = 0
	t.title = ""
	t.description = ""
	t.completed = false
	t.deleted = true
}

func (t *Task) CompleteTask() {
	t.completed = true
}

var taskId = 0
var tasksArray []Task
