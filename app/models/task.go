package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          uint           `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	IsDone      bool           `json:"is_done"`
	UserID      uint           `json:"user_id"`
	User        User           `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type PostTask struct {
	ID          uint           `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	IsDone      bool           `json:"is_done"`
	UserID      uint           `json:"user_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type CreateTaskInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func StoreTask(taskInput CreateTaskInput) PostTask {
	task := PostTask{Name: taskInput.Name, Description: taskInput.Description, IsDone: false, UserID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}

	DB.Create(&task)

	return task
}

func GetTasks() []Task {
	var tasks []Task

	DB.Preload("User").Find(&tasks)

	return tasks
}

func ShowTask(id string) Task {
	var task Task

	DB.Preload("User").First(&task, id)

	return task
}
