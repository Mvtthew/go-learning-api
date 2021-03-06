package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"strconv"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Token    string
}

func initializeUserMigration() {
	db := initDb()
	defer db.Close()

	db.AutoMigrate(&User{})
}

func GetUser(w http.ResponseWriter, r *http.Request)  {
	db := initDb()
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	password := vars["password"]

	var user User
	db.Where("name = ?", name).Where("password = ?", password).Find(&user)

	if user.ID != 0 {
		json.NewEncoder(w).Encode(user)
		return
	}

	response := make(map[string]string)
	response["message"] = "Bad credentials"

	json.NewEncoder(w).Encode(response)
}

func AllUsers(w http.ResponseWriter, r *http.Request) {

	_, err := checkToken(r)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	db := initDb()
	defer db.Close()

	var users []User
	db.Find(&users)

	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	db := initDb()
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})

	charset := "abcdefghijkmnopqrstuvwxyz0123456789"
	var password, token string
	for i := 0; i < 4; i++ {
		password += string(charset[rand.Intn(len(charset))])
	}
	for i := 0; i < 32; i++ {
		token += string(charset[rand.Intn(len(charset))])
	}

	var user User
	db.Last(&user)
	user.Password = strconv.Itoa(int(user.ID)) + password
	user.Token = strconv.Itoa(int(user.ID)) + token
	db.Save(&user)

	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	_, err := checkToken(r)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	db := initDb()
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]
	response := make(map[string]string)

	var user User
	db.Where("id = ?", id).First(&user)

	if user.ID != 0 {
		json.NewEncoder(w).Encode(user)
		db.Delete(&user)

		response["message"] = "User successfully deleted!"

	} else {
		response["message"] = "Could not find any user with this id!"
	}

	json.NewEncoder(w).Encode(response)

}
