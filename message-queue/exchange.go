package main

type Exchange struct {
	Name     string
	Bindings map[string][]*Queue //A map that holds the routing key and the corresponding queue that are bound to that routing key.
}
