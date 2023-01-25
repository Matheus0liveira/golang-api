package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/user", getUser)
	fmt.Println("api is on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
}

func getUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

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
