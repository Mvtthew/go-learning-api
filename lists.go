package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type List struct {
	gorm.Model
	UserId uint
	Name   string
}

func initializeListMigration() {
	db := initDb()
	defer db.Close()

	db.AutoMigrate(&List{})
}

func GetMyLists(w http.ResponseWriter, r *http.Request) {

	db := initDb()
	defer db.Close()

	user, err := checkToken(r)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	var userList []List
	db.Where("user_id = ?", user.ID).Find(&userList)

	json.NewEncoder(w).Encode(userList)

}

func CreateNewList(w http.ResponseWriter, r *http.Request) {

	db := initDb()
	defer db.Close()

	user, err := checkToken(r)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	vars := mux.Vars(r)
	name := vars["name"]

	db.Create(&List{Name: name, UserId: user.ID})

	response := make(map[string]string)
	response["message"] = "List created successfully!"
	json.NewEncoder(w).Encode(response)
}
