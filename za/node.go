package za

import (
	"fmt"
	"github.com/google/uuid"
)

type Node struct {
	id string // uuid
	isMiner bool
	localIp string 
	knownNodes []string // ipAddresses of known nodes
}


func createNode() *Node{
	return &Node{id:uuid.New().String()}
}

func startNode(node *Node){
	fmt.Printf("Node %s is running\n",node.id)
}

