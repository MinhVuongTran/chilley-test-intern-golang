package model

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var tasks = make(map[int]Task)
var nextID = 1

func GetNextID() int {
	return nextID
}

func IncrementNextID() {
	nextID++
}

func GetTasks() map[int]Task {
	return tasks
}
