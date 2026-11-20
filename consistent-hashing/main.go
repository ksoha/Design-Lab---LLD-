package main

import "fmt"

func main() {

	// Create a hash ring with 3 virtual nodes per unit of weight
	ring := NewHashRing(3)

	// Add three physical servers
	ring.AddNode(Node{ID: "server-A", Weight: 1})
	ring.AddNode(Node{ID: "server-B", Weight: 2})
	ring.AddNode(Node{ID: "server-C", Weight: 1})

	// Requests to route
	keys := []string{
		"user:101",
		"user:102",
		"user:103",
		"user:104",
		"user:105",
	}

	fmt.Println("Before removing server-B:")

	for _, key := range keys {
		server, err := ring.GetNode(key)

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("%s → %s\n", key, server)
	}

	// Remove server-B
	ring.RemoveNode("server-B")

	fmt.Println("\nAfter removing server-B:")

	for _, key := range keys {
		server, err := ring.GetNode(key)

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("%s → %s\n", key, server)
	}

	// Verify server-B has no remaining virtual positions
	for _, serverID := range ring.nodeMap {
		if serverID == "server-B" {
			panic("server-B was not completely removed")
		}
	}

	fmt.Println("\nAll tests passed!")
}
