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

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", hello).Methods("GET")

	initializeMigration()

	log.Fatal(http.ListenAndServe(":8082", router))

}
