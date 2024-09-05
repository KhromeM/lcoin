package za


// Classes:
type Node struct {
	Address string
	Blockchain *Blockchain
	Peers map[string]bool 
}
type Message struct {
	Type string
	Payload []byte
}



// Node methods:

func NewNode(address string) *Node {
	return &Node {
		Address: address,
		Blockchain: NewBlockchain(), // shouldn't do this and get the longest blockchain from peers instead
		Peers: make(map[string]bool), // should have some hardcoded peers
	}
}

func (n *Node) AddPeer(address string) {
	n.Peers[address] = true
}

func (n *Node) BrodcastBlock(b *Block) {
	for peer := range n.Peers {
		go n.sendBlock(peer, block)
	}
}

func (n *Node) sendBlock(peer string, b *block) {
	conn, err := net.dial("tcp", peer)
	if err != nil {
		log.Printf("Failed to connect to peer %s: %v", peer, err)
		return
	}
	defer conn.Close()

	encoder := gob.NewEncoder(conn) 
	message := Message(Type: "block", Payload: n.serializeBlock(block))
	err = encoder.Encode(message) 
	if err != nil {
		log.Printf("Failed to send block to peer %s: %v", peer, err)
	}
}

func (n *Node) serializeBlock(b *Block) []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(block)
	if err != nil {
		log.Printf("Failed to encode block. %v", err)
	}
	return result.Bytes()
}

func (n *Node) deserializeBlock(b []bytes) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(b))
	err := decoder.Decode(&block)
	if err != nil {
		log.Printf("Failed to decode block %v", err)
	}
	return &block
}


