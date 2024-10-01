package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

func calculateHash(index int, timestamp string, amount float64, prevHash string, nonce int) string {
	record := strconv.Itoa(index) + timestamp + fmt.Sprintf("%f", amount) + prevHash + strconv.Itoa(nonce)
	hash := sha1.New()
	hash.Write([]byte(record))
	return hex.EncodeToString(hash.Sum(nil))
}

func (bc *Blockchain) addGenesisBlock() {
	genesisBlock := Block{0, time.Now().String(), 0, "0", "", 0}
	genesisBlock.Hash = calculateHash(genesisBlock.Index, genesisBlock.Timestamp, genesisBlock.Amount, genesisBlock.PrevHash, genesisBlock.Nonce)
	bc.Chain = append(bc.Chain, genesisBlock)
	fmt.Println("Added the first Block i.e Genesis Block")
}

func (bc *Blockchain) addBlock(newBlock Block) {
	newBlock.PrevHash = bc.Chain[len(bc.Chain)-1].Hash
	newBlock.Hash = calculateHash(newBlock.Index, newBlock.Timestamp, newBlock.Amount, newBlock.PrevHash, newBlock.Nonce)
	bc.Chain = append(bc.Chain, newBlock)
}

// New Block with transaction data
// func (bc *Blockchain) mineBlock(nonce int, previousBlock Block) Block {
// 	for {
// 		hash := calculateHash(previousBlock.Index+1, time.Now().String(), 0, previousBlock.Hash, nonce)
// 		fmt.Println(hash[:bc.Difficulty])
// 		fmt.Println(string(make([]byte, bc.Difficulty)))
// 		if hash[:bc.Difficulty] == string(make([]byte, bc.Difficulty)) {

// 			fmt.Println("Mine Block is implemented")
// 			return Block{
// 				Index:     previousBlock.Index + 1,
// 				Timestamp: time.Now().String(),
// 				Amount:    0,
// 				Hash:      hash,
// 				PrevHash:  previousBlock.Hash,
// 				Nonce:     nonce,
// 			}
// 		}
// 		nonce++
// 	}
// }

// Mine a new block with a maximum number of attempts and a time limit
func (bc *Blockchain) mineBlock(maxAttempts int, timeLimit time.Duration, previousBlock Block) (Block, bool) {

	// Create a string of leading zeros based on difficulty
	target := ""
	for i := 0; i < bc.Difficulty; i++ {
		target += "0"
	}

	// ecdsa.
	// 	sa

	for attempts := 0; attempts < maxAttempts; attempts++ {
		nonce := attempts
		hash := calculateHash(previousBlock.Index+1, time.Now().String(), 0, previousBlock.Hash, nonce)
		// Check if the hash meets the difficulty requirement
		fmt.Println("all hash", hash[:bc.Difficulty])
		if hash[:bc.Difficulty] == target {
			fmt.Println("valid hash>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", hash[:bc.Difficulty])
			return Block{
				Index:     previousBlock.Index + 1,
				Timestamp: time.Now().String(),
				Amount:    0,
				Hash:      hash,
				PrevHash:  previousBlock.Hash,
				Nonce:     nonce,
			}, true
		}

		// Check if the time limit has been exceeded
		// if time.Since(startTime) > timeLimit {
		// 	break
		// }
	}

	// Return a zero-value block and false to indicate failure
	return Block{}, false
}

func (bc *Blockchain) createTransaction(transactionID string, amount float64) {
	if bc.validateTransaction(transactionID, amount) {
		fmt.Println("Transaction is validated")
		bc.TransactionPool[transactionID] = amount
	}
}

func (bc Blockchain) validateTransaction(transactionID string, amount float64) bool {
	// Ensure the amount is positive
	if amount <= 0 {
		fmt.Println("Invalid transaction amount")
		return false
	}

	// Check if transaction ID is already in the pool
	if _, exists := bc.TransactionPool[transactionID]; exists {
		fmt.Println("Duplicate transaction ID")
		return false
	}

	// If all checks pass, add the transaction ID to the pool
	bc.TransactionPool[transactionID] = amount
	return true
}

// func (bc *Blockchain) finalizeTransactions() {
// 	for transactionID, amount := range bc.TransactionPool {
// 		fmt.Println("Transaction ID", transactionID)
// 		newBlock := bc.mineBlock(0, bc.Chain[len(bc.Chain)-1])
// 		newBlock.Amount = amount
// 		bc.addBlock(newBlock)
// 		delete(bc.TransactionPool, transactionID) // Remove transaction from pool after adding to blockchain
// 	}
// }

func (bc *Blockchain) finalizeTransactions() {
	// Example parameters for mining
	maxAttempts := 1000000
	timeLimit := 5 * time.Second

	fmt.Println("Pool ", bc.TransactionPool)

	for transactionID, amount := range bc.TransactionPool {
		previousBlock := bc.Chain[len(bc.Chain)-1]
		newBlock, success := bc.mineBlock(maxAttempts, timeLimit, previousBlock)

		if success {
			newBlock.Amount = amount
			bc.addBlock(newBlock)
			delete(bc.TransactionPool, transactionID) // Remove transaction from pool after adding to blockchain
		} else {
			fmt.Println("Failed to mine block for transaction", transactionID)
		}
	}
}
