package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	userName string `json:"username" gorm:"unique"`
	passsword string `json:"password"`
	todos []Todo `gorm:"foregnKey:userID"`
}