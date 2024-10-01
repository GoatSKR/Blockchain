package main

import "fmt"

// My Block chain single block structure
type Block struct {
	Index     int
	Timestamp string
	Amount    float64
	Hash      string
	PrevHash  string
	Nonce     int
}

// Blockchain system structure
type Blockchain struct {
	Chain           []Block
	Difficulty      int
	TransactionPool map[string]float64 // Map to hold unique transactions
}

func NewBlockchain() *Blockchain {
	bc := &Blockchain{
		Chain:           []Block{},
		Difficulty:      4,                        // Number of leading zeros required in hash (adjust as needed)
		TransactionPool: make(map[string]float64), // Initialize map for transaction pool
	}
	fmt.Println("Added new blockchain")
	bc.addGenesisBlock()
	return bc
}

func main() {
	bc := NewBlockchain()

	// Create and validate transactions
	bc.createTransaction("tx1", 10.0)
	bc.createTransaction("tx2", 20.5)

	// Mine and add transactions to the blockchain
	bc.finalizeTransactions()

	// Display the blockchain
	for _, block := range bc.Chain {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Amount: %f\n", block.Amount)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Println()
	}
}
