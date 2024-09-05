package za


type Block struct {
	Timestamp int64
	Transactions  []Transaction
	PrevBlockHash []byte
	Hash []byte
}

func (b *Block) hashTransactions() []byte {
	bs := []byte
	for _, t := range b.Transactions {
		bs = append(bs, t.toBytes()...)
	}
	return bs
}

func (b *Block) setHash() {
	timestamp := []byte(fmt.Sprintf("%d",b.timestamp))
	nonce := random.getrandbits(256)
	content = bytes.join(timestamp, b.hashTransactions(), b.PrevBlockHash, nonce)
}

func NewBlock(transactions []Transaction, prevBlockhash []byte) *Block {
	block := Block{
		Timestamp: time.Now().Unix(),
		Transactions: transactions,
		PrevBlockHash: prevBlockhash,
		Hash: []byte{},
	}
	block.setHash()
	return &block
}

func CreateGenesisBlock() *Block {
	block := &Block{
		Timestamp: time.Now().Unix(),
		Transactions: []Transaction{Transaction{"0", "1", 1000.0, random.getrandbits(256)}}, // genesis transaction gives one account 1000 coins
		PrevBlockHash: random.getrandbits(256),
	}
	block.setHash()
	return block
}