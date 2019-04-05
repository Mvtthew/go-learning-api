package main

import "github.com/jinzhu/gorm"

func initDb()*gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to database!")
	}
	return db
}
