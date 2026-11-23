package payment

type Payment struct {
	ID      int
	OrderID int
	Amount  float64
	Status  string
}

var payments []Payment //slice to store payments

func CreatePayment(orderID int, amount float64) Payment {
	payment := Payment{
		ID:      len(payments) + 1,
		OrderID: orderID,
		Amount:  amount,
		Status:  "SUCCESS",
	}

	payments = append(payments, payment)

	return payment
}
