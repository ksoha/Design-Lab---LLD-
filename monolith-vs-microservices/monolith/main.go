package main

import (
	"encoding/json"
	"net/http"

	"github.com/ksoha/Design-Lab---LLD-/monolith-vs-microservices/monolith/order"
	"github.com/ksoha/Design-Lab---LLD-/monolith-vs-microservices/monolith/users"
)

func main() {
	http.HandleFunc("/users", createUserHandler)
	http.HandleFunc("/orders", createOrderHandler)

	http.ListenAndServe(":8080", nil)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	json.NewDecoder(r.Body).Decode(&input)

	newUser := users.CreateUser(input.Name, input.Email)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func createOrderHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		UserID int     `json:"user_id"`
		Amount float64 `json:"amount"`
	}

	json.NewDecoder(r.Body).Decode(&input)

	newOrder := order.CreateOrder(input.UserID, input.Amount)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newOrder)
}
