# BlockchainInGo
Just a simple blockchain core in golang for education purposes
- Use SHA-256 :), for computing the hash of the blocks

## Proof of Work
- Use a consensus algorithm to compute nonce values 
- `` nonce from 1 to infinity + previous hash + transaction ``
- Compute hash of previousHash, transactions and nonce.
- Based on the ``MINING_DIFFICULTY`` value. The number of zeroes at the beginning of the hash
gives us the nonce values which is the rounds you need to compute hashes till get those zeros at the
beginning of the final hash.
- 
```go
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

```
