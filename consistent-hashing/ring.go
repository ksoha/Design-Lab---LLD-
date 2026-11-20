package main

import (
	"fmt"
	"hash/fnv"
	"sort"
)

//HashRing represents a consistent hashing ring
//nodes : store all hash positions on the ring in a sorted order
//nodeMap : maps each hash position to its physical server id
//pointMultipier : Controls how many virtual positions each physical server receives

type HashRing struct {
	nodes           []uint64
	nodeMap         map[uint64]string //key -> postion, value -> node id
	pointMultiplier int
}

// hash function
func hashKey(key string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(key)) //converting string in bytes
	return h.Sum64()     //returning 64 bit hash value
}

// contsructor for Ring
func NewHashRing(pointMultiplier int) *HashRing {
	return &HashRing{
		nodes:           []uint64{},
		nodeMap:         make(map[uint64]string),
		pointMultiplier: pointMultiplier,
	}
}

// AddNode adds a physical server to the ring with virtual nodes
func (r *HashRing) AddNode(node Node) {
	//calculating number of virtual positions
	totalPoints := r.pointMultiplier * node.Weight

	for i := 0; i < totalPoints; i++ {

		//creating virtual node id
		virtualNode := fmt.Sprintf("%s#%d", node.ID, i)

		position := hashKey(virtualNode) //hsaing the virtual servers(nodes)

		r.nodes = append(r.nodes, position)
		r.nodeMap[position] = node.ID
	}

	//sorting the nodes slice to maintain order on the ring
	sort.Slice(r.nodes, func(i, j int) bool {
		return r.nodes[i] < r.nodes[j]
	})
}

// GetNode returns the physical server id for a given key
func (r *HashRing) GetNode(key string) (string, error) {

	//check whether the ring is empty \
	if len(r.nodes) == 0 {
		return "", fmt.Errorf("Hash ring is empty")
	}

	position := hashKey(key) //hashing the key

	//finding the next clockwise node position on the ring
	idx := sort.Search(len(r.nodes), func(i int) bool {
		return r.nodes[i] >= position
	})

	if idx == len(r.nodes) {
		idx = 0 //wrap around to the first node if we reach end of the ring
	}

	//finding the physical server id for the found position
	serverPosition := r.nodes[idx]

	return r.nodeMap[serverPosition], nil

}
