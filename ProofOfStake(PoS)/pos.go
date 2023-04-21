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
}

type Validator struct {
	Address string
	Balance int
}

type Blockchain struct {
	Blocks     []*Block
	Validators []*Validator
}

func NewBlock(index int, data string, prevHash string) *Block {
	block := &Block{
		Index:     index,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevHash,
	}
	block.Hash = block.calculateHash()
	return block
}

func (b *Block) calculateHash() string {
	record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func NewValidator(address string, balance int) *Validator {
	return &Validator{
		Address: address,
		Balance: balance,
	}
}

func NewBlockchain() *Blockchain {
	block := NewBlock(0, "Genesis Block", "")
	validators := []*Validator{
		NewValidator("Address 1", 100),
		NewValidator("Address 2", 200),
		NewValidator("Address 3", 300),
	}
	return &Blockchain{
		Blocks:     []*Block{block},
		Validators: validators,
	}
}

func (bc *Blockchain) AddBlock(data string, validator *Validator) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(prevBlock.Index+1, data, prevBlock.Hash)
	newBlock.Hash = newBlock.calculateHash()
	bc.Blocks = append(bc.Blocks, newBlock)
	validator.Balance += 1
}

func (bc *Blockchain) ChooseValidator() *Validator {
	totalBalance := 0
	for _, v := range bc.Validators {
		totalBalance += v.Balance
	}
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(totalBalance)
	for _, v := range bc.Validators {
		if r < v.Balance {
			return v
		}
		r -= v.Balance
	}
	return nil
}

func main() {
	bc := NewBlockchain()

	for i := 0; i < 10; i++ {
		validator := bc.ChooseValidator()
		bc.AddBlock(fmt.Sprintf("Block %d Data", i+1), validator)
		fmt.Printf("Block %d mined by validator %s\n", i+1, validator.Address)
	}

	for _, block := range bc.Blocks {
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
	}
}
