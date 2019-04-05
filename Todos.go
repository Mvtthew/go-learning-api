package main

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

type Todo struct {
	gorm.Model
	UserId uint
	ListId uint
	Name   string
	Done   bool
}

func initializeTodoMigration() {

	db := initDb()
	defer db.Close()

	db.AutoMigrate(&Todo{})

}

func GetMyTodos(w http.ResponseWriter, r *http.Request) {

	db := initDb()
	defer db.Close()

	user, err := checkToken(r)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	vars := mux.Vars(r)
	listId := vars["listId"]

	var todos []Todo
	db.Where("user_id = ?", user.ID).Where("list_id = ?", listId).Find(&todos)

	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {

	db := initDb()
	defer db.Close()

	user, err := checkToken(r)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	vars := mux.Vars(r)
	listId, _ := strconv.Atoi(vars["listId"])
	name := vars["name"]

	db.Create(&Todo{UserId: user.ID, ListId: uint(listId), Name: name})

	response := make(map[string]string)
	response["message"] = "Todo created successfully!"

	json.NewEncoder(w).Encode(response)
}
