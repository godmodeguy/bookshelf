package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open("sqlite3", "storage/test.db")

	if err != nil {
		panic("can't connect to db")
	}

	db.AutoMigrate(&Book{})

	DB = db
}
