package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dataBase *gorm.DB

func connectDB()  {
	var err error
	dataBase,err = gorm.Open(sqlite.Open("todos.db"))

	if err != nil {
		log.Fatal("Connecting to db : Failed!")
	}
}