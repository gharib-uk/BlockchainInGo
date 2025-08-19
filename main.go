package main

import (
	"BlockchainInGo/block"
	"BlockchainInGo/wallet"
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	myBlockchainAddress := "my_blockchain_address"
	blockChain := block.NewBlockchain(myBlockchainAddress)
	println(blockChain.ToString())

	blockChain.AddTransaction("A", "B", 1.0)
	previousHash := blockChain.LastBlock().Hash()
	nonce := blockChain.ProofOfWork()
	blockChain.CreateBlock(nonce, previousHash)
	println(blockChain.ToString())

	blockChain.AddTransaction("C", "D", 2.0)
	blockChain.AddTransaction("X", "Y", 3.0)
	previousHash = blockChain.LastBlock().Hash()
	nonce = blockChain.ProofOfWork()
	blockChain.CreateBlock(nonce, previousHash)
	println(blockChain.ToString())

	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())
}
