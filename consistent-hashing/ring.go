package main

//Ring represents a consistent hashing ring
//nodes : store all hash positions on the ring in a sorted order
//nodeMap : maps each hash position to its physical server id
//pointMultipier : Controls how many virtual positions each physical server receives

type Ring struct {
	nodes           []uint64
	nodeMap         map[uint64]string
	pointMultiplier int
}
