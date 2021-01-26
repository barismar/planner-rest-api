package models

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

type Task struct {
	gorm.Model
	ID          int    `"json:id"`
	Description string `json:"description"`
	IsDone      bool   `json:"isDone"`
}

func GetTasks() []Task {
	var tasks []Task

	DB.Find(&tasks)

	return tasks
}

func ShowTask(id string) Task {
	var task Task

	DB.First(&task, id)

	return task
}
