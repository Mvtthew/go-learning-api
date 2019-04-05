package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"log"
)

func hello(w http.ResponseWriter, r *http.Request)  {
	 fmt.Fprintf(w, "/ endpoiont")
}

func handleFunctions()  {
	router := mux.NewRouter()

	router.HandleFunc("/", hello).Methods("GET")
	router.HandleFunc("/users", AllUsers).Methods("GET")

	router.HandleFunc("/user/{name}/{email}", NewUser).Methods("POST")
	router.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")

	fmt.Println("REST Server up on port :8082")
	log.Fatal(http.ListenAndServe(":8082", router))
}

func main() {

	initializeMigration()
	handleFunctions()

}
