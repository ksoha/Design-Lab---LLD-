package order

import payment "github.com/ksoha/Design-Lab---LLD-/monolith-vs-microservices/monolith/payments"

type Order struct {
	ID        int
	UserID    int
	Amount    float64
	Status    string
	PaymentID int
}

var orders []Order

func CreateOrder(userID int, amount float64) Order {
	order := Order{
		ID:     len(orders) + 1,
		UserID: userID,
		Amount: amount,
		Status: "CREATED",
	}

	payment := payment.CreatePayment(order.ID, amount)

	order.PaymentID = payment.ID
	order.Status = "PAID"

	orders = append(orders, order)

	return order
}
