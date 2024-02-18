package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Function to test get items API
func TestGetItems(t *testing.T) {
	req, err := http.NewRequest("GET", "/items", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(itemsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Item handler returned incorrect status code: got %v, expected %v", status, http.StatusOK)
	}
}

// Function to test add items API
func TestAddItem(t *testing.T) {
	item := Item{Name: "New Item"}
	body, _ := json.Marshal(item)
	req, err := http.NewRequest("POST", "/items", bytes.NewBuffer(body))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(itemsHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Item handler returned incorrect status code: got %d, expected %d", http.StatusCreated, rr.Code)
	}

	var newItem Item
	json.Unmarshal(rr.Body.Bytes(), &newItem)
	if newItem.Name != item.Name {
		t.Errorf("Item handler returned incorrect item name: got %s, expected %s", newItem.Name, item.Name)
	}
}
