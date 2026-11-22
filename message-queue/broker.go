package main

// a broker creates queues and exchanges
type Broker struct {
	Queues    map[string]*Queue    //a map of queue names to their corresponding Queue objects
	Exchanges map[string]*Exchange //a map of exchange names to their corresponding Exchange objects
}

func NewBroker() *Broker {
	return &Broker{
		Queues:    make(map[string]*Queue),    //initializing the Queues map
		Exchanges: make(map[string]*Exchange), //initializing the Exchanges map
	}
}

func (b *Broker) CreateQueue(name string) *Queue {
	queue := NewQueue(name)

	b.Queues[name] = queue //adding the newly created queue to the Queues map

	return queue
}

func (b *Broker) CreateExchange(name string) *Exchange {
	exchange := NewExchange(name)

	b.Exchanges[name] = exchange //adding the newly created exchange to the Exchanges map

	return exchange
}
