package main

import (
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

type Block struct {
	nonce        int
	previousHash [64]byte
	timestamp    int64
	transactions []string
}

type Blockchain struct {
	chain           []*Block
	transactionPool []*Transaction
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

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
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

func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}

func (t *Transaction) ToString() string {

	return fmt.Sprintf(
		"%s\n sender_blockchain_address      %s\n recipient_blockchain_address   %s\n"+
			" value                          %.1f\n",
		strings.Repeat("-", 40), t.senderBlockchainAddress, t.recipientBlockchainAddress, t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
	})
}
