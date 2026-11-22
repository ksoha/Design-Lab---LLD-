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
