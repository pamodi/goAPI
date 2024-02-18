package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// Define Item struct
type Item struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
}

// In-memory storage for items
var items = []Item{}

const Dport = ":8012"

// Items handler to handle get and post requests for items
func itemsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getItems(w, r)
	case http.MethodPost:
		addItems(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// Function to get all the items
func getItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(items)
}

// Function to add an item
func addItems(w http.ResponseWriter, r *http.Request) {
	var newItem Item

	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate a unique ID for the item
	newItem.ID = uuid.New().String()
	items = append(items, newItem)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
}

func main() {
	http.HandleFunc("/items", itemsHandler)
	http.HandleFunc("/items/", itemsHandler)

	fmt.Printf("Server is starting on port: %v\n", Dport)
	http.ListenAndServe(":8012", nil)
}
