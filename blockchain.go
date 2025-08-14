package main

import (
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	nonce        int
	previousHash [64]byte
	timestamp    int64
	transactions []string
}

type Blockchain struct {
	chain           []*Block
	transactionPool []string
}

func NewBlockchain() *Blockchain {
	b := new(Block)
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

func NewBlock(nonce int, previousHash [64]byte) *Block {
	return &Block{
		nonce:        nonce,
		previousHash: previousHash,
		timestamp:    time.Now().UnixNano(),
	}
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [64]byte) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) ToString() string {
	str := "Blockchain: \n"
	for _, b := range bc.chain {
		str += b.ToString() + "\n"
	}

	return str
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (block *Block) ToString() string {
	// with block.previousHash[:] create a byte slice
	return fmt.Sprintf("Block:\ntimestamp       %d\nnonce           %d\nprevious_hash   %x\ntransactions    %s\n",
		block.timestamp, block.nonce, block.previousHash, block.transactions)
}

func (block *Block) Hash() [64]byte {
	m, _ := json.Marshal(block)
	return sha512.Sum512([]byte(m))
}

func (block *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64    `json:"timestamp"`
		Nonce        int      `json:"nonce"`
		PreviousHash [64]byte `json:"previous_hash"`
		Transactions []string `json:"transactions"`
	}{
		Timestamp:    block.timestamp,
		Nonce:        block.nonce,
		PreviousHash: block.previousHash,
		Transactions: block.transactions,
	})
}
