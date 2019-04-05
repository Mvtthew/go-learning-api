package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Name string
	Email string
}

func initializeMigration() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect do database!")
	}
	defer db.Close()

	db.AutoMigrate(&User{})

}
