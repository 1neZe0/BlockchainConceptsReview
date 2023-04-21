package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

const difficulty = 2

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
	Validator string
}

type Blockchain struct {
	Blocks     []*Block
	Validators []string
}

func NewBlock(index int, data string, prevHash string, validator string) *Block {
	block := &Block{
		Index:     index,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevHash,
		Validator: validator,
	}
	block.Hash = block.calculateHash()
	return block
}

func (b *Block) calculateHash() string {
	record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PrevHash + b.Validator
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func NewBlockchain() *Blockchain {
	block := NewBlock(0, "Genesis Block", "", "")
	validators := []string{
		"Validator 1",
		"Validator 2",
		"Validator 3",
	}
	return &Blockchain{
		Blocks:     []*Block{block},
		Validators: validators,
	}
}

func (bc *Blockchain) AddBlock(data string, validator string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(prevBlock.Index+1, data, prevBlock.Hash, validator)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) GetLastBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

func (bc *Blockchain) ValidateBlock(block *Block) bool {
	if block.Index == 0 {
		return true
	}
	prevBlock := bc.Blocks[block.Index-1]
	if block.PrevHash != prevBlock.Hash {
		return false
	}
	if block.Validator != prevBlock.Validator {
		return false
	}
	if block.calculateHash() != block.Hash {
		return false
	}
	return true
}

func (bc *Blockchain) Consensus() {
	// Count the number of blocks validated by each validator
	blockCount := make(map[string]int)
	for _, block := range bc.Blocks {
		blockCount[block.Validator]++
	}

	// Find the validator with the most validated blocks
	var validator string
	maxCount := 0
	for _, v := range bc.Validators {
		if blockCount[v] > maxCount {
			maxCount = blockCount[v]
			validator = v
		}
	}

	// Replace the blockchain with the longest validated chain
	newChain := make([]*Block, 0)
	for _, block := range bc.Blocks {
		if block.Validator == validator {
			newChain = append(newChain, block)
		}
	}
	bc.Blocks = newChain
}

func main() {
	bc := NewBlockchain()

	fmt.Println("Mining block 1...")
	bc.AddBlock("Block 1 Data", "Validator 1")

	fmt.Println("Mining block 2...")
	bc.AddBlock("Block 2 Data", "Validator 2")

	fmt.Println("Mining block 3...")
	bc.AddBlock("Block 3 Data", "Validator 1")

	fmt.Println("Mining block 4...")
	bc.AddBlock("Block 4 Data", "Validator 3")

	fmt.Println("Blockchain:")
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Validator: %s\n", block.Validator)
		fmt.Println()
	}

	// Validate the blockchain
	fmt.Println("Validating blockchain...")
	for _, block := range bc.Blocks {
		if !bc.ValidateBlock(block) {
			fmt.Printf("Invalid block: %d\n", block.Index)
			return
		}
	}

	// Perform consensus
	fmt.Println("Performing consensus...")
	bc.Consensus()

	// Print the validated blockchain
	fmt.Println("Validated blockchain:")
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Validator: %s\n", block.Validator)
		fmt.Println()
	}
}
