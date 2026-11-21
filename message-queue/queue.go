package main

type Queue struct {
	Name    string
	Message []Message //A slice containing multiple message objects.
}

// Constructor
func NewQueue(name string) *Queue {
	return &Queue{
		Name:    name,
		Message: []Message{}, //intializing the empty slice so that the queue is ready to accecpt messages
	}
}

// function that adds a message to the queue
func (q *Queue) Enqueue(msg Message) {
	q.Message = append(q.Message, msg)
}

// functions that removes a message from the queue
func (q *Queue) Dequeue() (Message, bool) {
	//in case the queue is empty
	if len(q.Message) == 0 {
		return Message{}, false
	}

	//removing the first message from the queue
	msg := q.Message[0]
	q.Message = q.Message[1:] //the queue now has messages fron idx 1 onwards

	return msg, true
}
