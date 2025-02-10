package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DataBase *gorm.DB

func connectDB()  {
	var err error
	DataBase,err = gorm.Open(sqlite.Open("todos.db"))

	if err != nil {
		log.Fatal("Connecting to db : Failed!")
	}
}