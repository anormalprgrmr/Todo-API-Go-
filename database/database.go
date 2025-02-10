package database

import (
	"Todo-API/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DataBase *gorm.DB

func ConnectDB()  {
	var err error
	DataBase,err = gorm.Open(sqlite.Open("todos.db"))

	if err != nil {
		log.Fatal("Connecting to db : Failed!")
	}

	DataBase.AutoMigrate(&models.User{}, &models.Todo{})
}