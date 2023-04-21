package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
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
	newBlock.Hash = newBlock.calculateHash()
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) ChooseValidator() string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(bc.Validators))
	return bc.Validators[r]
}

func main() {
	bc := NewBlockchain()

	for i := 0; i < 10; i++ {
		validator := bc.ChooseValidator()
		bc.AddBlock(fmt.Sprintf("Block %d Data", i+1), validator)
		fmt.Printf("Block %d mined by validator %s\n", i+1, validator)
	}

	for _, block := range bc.Blocks {
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Validator: %s\n", block.Validator)
	}
}
