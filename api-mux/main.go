package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	muxResponse := mux.NewRouter()

	muxResponse.HandleFunc("/user", getUser).Methods("GET")
	fmt.Println("api started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", muxResponse))
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
}

func getUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]User{{
		ID:       1,
		Name:     "Matheus",
		LastName: "Oliveira",
	},
		{
			ID:       1,
			Name:     "Matheus 1",
			LastName: "Oliveira 1",
		}})

	return
}
