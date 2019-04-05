package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"log"
)

func hello(w http.ResponseWriter, r *http.Request)  {
	 fmt.Fprintf(w, "Mvtthew GO api running! :)")
}

func handleFunctions()  {
	router := mux.NewRouter()

	router.HandleFunc("/", hello).Methods("GET")

	router.HandleFunc("/users", AllUsers).Methods("GET")
	router.HandleFunc("/login/{name}/{password}", GetUser).Methods("POST") // user data
	router.HandleFunc("/register/{name}/{email}", NewUser).Methods("POST") // creating new user
	router.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE") // deleting user

	router.HandleFunc("/my/lists", GetMyLists).Methods("GET")
	router.HandleFunc("/my/lists/{name}", CreateNewList).Methods("POST")

	router.HandleFunc("/my/todos/{listId}", GetMyTodos).Methods("GET")
	router.HandleFunc("/my/todos/{listId}/{name}", CreateTodo).Methods("POST")

	fmt.Println("REST Server up on port :8082")
	log.Fatal(http.ListenAndServe(":8082", router))
}

func main() {

	initializeUserMigration() // users
	initializeListMigration() // lists
	initializeTodoMigration() // todos
	handleFunctions()

}
