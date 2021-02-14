package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDBConnection() *gorm.DB {
	dataSourceName := Env("DB_USERNAME", "") + ":" + Env("DB_PASSWORD", "") + "@tcp(" + Env("DB_HOST", "127.0.0.1") + ":" + Env("DB_PORT", "3306") + ")/" + Env("DB_NAME", "") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
