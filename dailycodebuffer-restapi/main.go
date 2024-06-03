package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	InitialMigration()
	initializeRouter()
}