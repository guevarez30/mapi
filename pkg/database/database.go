package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	dsn := os.Getenv("DSN")
	con := dsn + "&parseTime=true"
	db, err := gorm.Open(mysql.Open(con), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
