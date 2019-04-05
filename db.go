package main

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"errors"
)

func initDb()*gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to database!")
	}
	return db
}

func checkToken(r *http.Request) (User, error) {
	token := r.Header.Get("token")

	var user User

	if token != "" {
		db := initDb()
		db.Where("token = ?", token).First(&user)
		if user.ID == 0 {
			// bad token response
			return user, errors.New("user unauthorized")
		} else {
			// user authorized
			return user, nil
		}
	}

	// request without token
	return user, errors.New("header 'token' required")
}