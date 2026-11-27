package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Order struct {
	ID        int     `json:"id"`
	UserID    int     `json:"user_id"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
	PaymentID int     `json:"payment_id"`
}

type PaymentResponse struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

var orders []Order

func createOrderHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		UserID int     `json:"user_id"`
		Amount float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	order := Order{
		ID:     len(orders) + 1,
		UserID: input.UserID,
		Amount: input.Amount,
		Status: "CREATED",
	}

	// Simulate a call to the payment service
	paymentRequest := map[string]interface{}{
		"order_id": order.ID,
		"amount":   order.Amount,
	}

	requestBody, err := json.Marshal(paymentRequest)
	if err != nil {
		http.Error(w, "Could not prepare payment request", http.StatusInternalServerError)
		return
	}

	resp, err := http.Post(
		"http://localhost:3003/payments",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		http.Error(w, "Payment service unavailable", http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		http.Error(w, "Payment failed", http.StatusBadGateway)
		return
	}

	var payment PaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&payment); err != nil {
		http.Error(w, "Invalid payment response", http.StatusBadGateway)
		return
	}

	order.PaymentID = payment.ID
	order.Status = "PAID"

	orders = append(orders, order)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func main() {
	http.HandleFunc("/orders", createOrderHandler)

	fmt.Println("Order Service running on port 3002")

	if err := http.ListenAndServe(":3002", nil); err != nil {
		panic(err)
	}
}
