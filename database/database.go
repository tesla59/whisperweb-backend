package database

import (
	"fmt"

	"github.com/tesla59/whisperweb-backend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func InitDB() {
	var err error
	DBConn, err = gorm.Open(sqlite.Open("confessions.db"))
	if err != nil {
		panic("Cannot connect to Database")
	}
	fmt.Println("Connected to DB")
	DBConn.AutoMigrate(&models.Confession{})
	fmt.Println("Migrated successfully")
}
