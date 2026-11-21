package main

type Message struct {
	RoutingKey string
	Body       string
}

//A message should ideally conatain :
//1.Routing Key: This is used to determine which queue the message should be sent to.
//2.Body: This is the actual content of the message.
