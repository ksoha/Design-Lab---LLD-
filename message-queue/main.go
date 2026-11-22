package main

import "fmt"

func main() {
	// 1. Create the broker
	broker := NewBroker()

	// 2. Create queues
	emailQueue := broker.CreateQueue("email_queue")
	smsQueue := broker.CreateQueue("sms_queue")
	inventoryQueue := broker.CreateQueue("inventory_queue")

	// 3. Create an exchange
	orderExchange := broker.CreateExchange("order_exchange")

	// 4. Bind queues to routing key
	orderExchange.BindQueue("order.created", emailQueue)
	orderExchange.BindQueue("order.created", smsQueue)
	orderExchange.BindQueue("order.created", inventoryQueue)

	// 5. Producer creates and publishes a message
	msg := Message{
		RoutingKey: "order.created",
		Body:       "Order ORD101 created",
	}

	orderExchange.Publish(msg)

	// 6. Consumers retrieve messages
	queues := []*Queue{
		emailQueue,
		smsQueue,
		inventoryQueue,
	}

	for _, queue := range queues {
		msg, ok := queue.Dequeue()

		if ok {
			fmt.Printf("%s processed: %s\n", queue.Name, msg.Body)
		}
	}
}
