package za

import "sync"


type Blockchain struct {
	Blocks []*Block
	mutex sync.Mutex
}


func NewBlockchain() *Blockchain {
	return &Blockchain{Blocks:[]*Block{CreateGenesisBlock()}, sync.Mutex{}}
}

func (bc *Blockchain) AddBlock(b *Block) {
	bc.mutex
}
