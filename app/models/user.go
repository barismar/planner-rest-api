package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id"`
	Email     string         `json:"email"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type CreateUserInput struct {
	Email string `json:"email" binding:"required"`
	Name  string `json:"name" binding:"required"`
}

func StoreUser(UserInput CreateUserInput) User {
	User := User{Email: UserInput.Email, Name: UserInput.Name, CreatedAt: time.Now(), UpdatedAt: time.Now()}

	DB.Create(&User)

	return User
}

func GetUsers() []User {
	var Users []User

	DB.Find(&Users)

	return Users
}

func ShowUser(id string) User {
	var User User

	DB.First(&User, id)

	return User
}
