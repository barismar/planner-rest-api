package models

type Task struct {
	ID          int    `"json:id"`
	Description string `json:"description"`
	IsDone      bool   `json:"isDone"`
}

func GetTasks() []Task {
	Tasks := []Task{
		Task{ID: 1, Description: "This is first task", IsDone: true},
		Task{ID: 2, Description: "This is second task", IsDone: false},
	}

	return Tasks
}

func ShowTask() Task {
	return Task{ID: 1, Description: "This is first task", IsDone: true}
}
