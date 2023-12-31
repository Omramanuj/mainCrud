package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()
	s := route.PathPrefix("/api").Subrouter() //Base Path

	//Routes

	s.HandleFunc("/CreateProfile", CreateProfile).Methods("POST")
	s.HandleFunc("/GetAllUsers", GetAllUsers).Methods("GET")
	s.HandleFunc("/GetUserProfile", GetUserProfile).Methods("POST")
	s.HandleFunc("/UpdateProfile", UpdateProfile).Methods("PUT")
	s.HandleFunc("/DeleteProfile/{id}", DeleteProfile).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", s)) // Run Server
}
