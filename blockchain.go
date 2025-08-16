package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const MINING_DIFFICULTY = 3

type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []*Transaction
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

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	b.transactions = transactions
	return b
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (block *Block) ToString() string {
	// with block.previousHash[:] create a byte slice
	s := fmt.Sprintf("Block:\ntimestamp       %d\nnonce           %d\nprevious_hash   %x\nTransactions:\n",
		block.timestamp, block.nonce, block.previousHash)

	for _, tx := range block.transactions {
		s += fmt.Sprintf("%s\n", tx.ToString())
	}

	return s
}

func (block *Block) Hash() [32]byte {
	m, _ := json.Marshal(block)
	return sha256.Sum256([]byte(m))
}

func (block *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		Nonce        int            `json:"nonce"`
		PreviousHash [32]byte       `json:"previous_hash"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Timestamp:    block.timestamp,
		Nonce:        block.nonce,
		PreviousHash: block.previousHash,
		Transactions: block.transactions,
	})
}

func (bc *Blockchain) ValidProof(nonce int, previousHash [32]byte,
	transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{
		nonce:        nonce,
		previousHash: previousHash,
		timestamp:    0,
		transactions: transactions,
	}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficulty] == zeros
}

func (bc *Blockchain) ProofOfWork() int {
	transactions := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().Hash()
	nonce := 0
	for !bc.ValidProof(nonce, previousHash, transactions, MINING_DIFFICULTY) {
		nonce += 1
	}
	return nonce
}

// CopyTransactionPool a hard copy function for transactionPool in Blockchain
func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, t := range bc.transactionPool {
		transactions = append(transactions,
			NewTransaction(t.senderBlockchainAddress,
				t.recipientBlockchainAddress,
				t.value))
	}
	return transactions
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
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
