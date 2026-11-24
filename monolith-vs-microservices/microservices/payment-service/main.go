package main

import (
	"encoding/json"
	"net/http"
)

type Payment struct {
	ID      int     `json:"id"`
	OrderID int     `json:"order_id"`
	Amount  float64 `json:"amount"`
	Status  string  `json:"status"`
}

var payments []Payment

func createPaymentHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		OrderID int     `json:"order_id"`
		Amount  float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	payment := Payment{
		ID:      len(payments) + 1,
		OrderID: input.OrderID,
		Amount:  input.Amount,
		Status:  "Pending",
	}

	payments = append(payments, payment)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(payment)
}

func main() {
	http.HandleFunc("/payments", createPaymentHandler)

	println("Payment service is running on port 3002")

	if err := http.ListenAndServe(":3003", nil); err != nil {
		panic(err)
	}
}
