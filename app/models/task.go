package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          uint           `json:"id"`
	Description string         `json:"description"`
	IsDone      bool           `json:"is_done"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

func StoreTask(description string) Task {
	task := Task{Description: description, IsDone: false, CreatedAt: time.Now(), UpdatedAt: time.Now()}

	DB.Create(&task)

	return task
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
