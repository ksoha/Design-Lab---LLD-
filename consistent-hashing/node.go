package main

// Node represents a physical server in the consisten hashing ring
type Node struct {
	ID     string
	Weight int
}
