package main

type Exchange struct {
	Name     string
	Bindings map[string][]*Queue //a slice of pointers to the Queues
}

func NewExchange(name string) *Exchange {
	return &Exchange{
		Name:     name,
		Bindings: make(map[string][]*Queue), //initializing the Binding map
	}
}

// function that binds a queue to the exchange with a specific routing key
func (e *Exchange) BindQueue(RoutingKey string, queue *Queue) {
	e.Bindings[RoutingKey] = append(e.Bindings[RoutingKey], queue) //adding a new queue to the existing slice
}

// Finds matching queues and enqueues the message into each
func (e *Exchange) Publish(msg Message) {
	//finding the queues that macth the routing key of the message
	queues := e.Bindings[msg.RoutingKey]

	//iterating over the slice of queues and enqueuing the message into each queue
	for _, queue := range queues {
		queue.Enqueue(msg)
	}
}
