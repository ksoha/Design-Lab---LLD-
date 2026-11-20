package main

import (
	"hash/fnv"
)

//Ring represents a consistent hashing ring
//nodes : store all hash positions on the ring in a sorted order
//nodeMap : maps each hash position to its physical server id
//pointMultipier : Controls how many virtual positions each physical server receives

type Ring struct {
	nodes           []uint64
	nodeMap         map[uint64]string
	pointMultiplier int
}

// hash function
func hashKey(key string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(key)) //converting string in bytes
	return h.Sum64()     //returning 64 bit hash value
}

// contsructor for Ring
func NewHashRing(pointMultiplier int) *Ring {
	return &Ring{
		nodes:           []uint64{},
		nodeMap:         make(map[uint64]string),
		pointMultiplier: pointMultiplier,
	}
}
