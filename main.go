package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

type Item struct {
    Name  string `json:"name"`
    Price int    `json:"price"`
}

var items = []Item{
    {"Aqua", 2000},
    {"Sosro", 5000},
    {"Cola", 7000},
    {"Milo", 9000},
    {"Coffee", 12000},
}


// Get all items
func getItems(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(items)
}

// Get single item by name
func getItem(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for _, item := range items {
        if item.Name == params["name"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode("Item not found")
}

// Create a new item
func createItem(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var item Item
    _ = json.NewDecoder(r.Body).Decode(&item)

    items = append(items, item)
    json.NewEncoder(w).Encode(item)
}

// Update an item
func updateItem(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    var updatedItem Item
    _ = json.NewDecoder(r.Body).Decode(&updatedItem)

    for i, item := range items {
        if item.Name == params["name"] {
            items[i] = updatedItem
            json.NewEncoder(w).Encode(updatedItem)
            return
        }
    }
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode("Item not found")
}

// Delete an item
func deleteItem(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for i, item := range items {
        if item.Name == params["name"] {
            items = append(items[:i], items[i+1:]...)
            json.NewEncoder(w).Encode("Item deleted")
            return
        }
    }
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode("Item not found")
}

func main() {
    router := mux.NewRouter()

    // Define routes
    router.HandleFunc("/items", getItems).Methods("GET")
    router.HandleFunc("/items/{name}", getItem).Methods("GET")
    router.HandleFunc("/items", createItem).Methods("POST")
    router.HandleFunc("/items/{name}", updateItem).Methods("PUT")
    router.HandleFunc("/items/{name}", deleteItem).Methods("DELETE")

    fmt.Println("Server starting on port 8000...")
    http.ListenAndServe(":8000", router)
}
