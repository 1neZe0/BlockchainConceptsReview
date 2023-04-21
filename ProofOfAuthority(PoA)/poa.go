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
	Data      int
	PrevHash  string
	Hash      string
	Validator string
}

type Blockchain struct {
	Blocks     []*Block
	Validators []Validator
}

type Validator struct {
	Name    string
	Respect int
}

func NewBlock(index int, data int, prevHash string, validator string) *Block {
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
	record := strconv.Itoa(b.Index) + b.Timestamp + strconv.Itoa(b.Data) + b.PrevHash + b.Validator
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func NewBlockchain() *Blockchain {
	block := NewBlock(0, 1, "", "")
	validators := []Validator{
		{"John Cena", 1_000_000},
		{"Bald From Brazzers", 10_000_000},
		{"Swarztransnigger", 100_000_000},
	}
	return &Blockchain{
		Blocks:     []*Block{block},
		Validators: validators,
	}
}

func (bc *Blockchain) AddBlock(data int, validator Validator) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(prevBlock.Index+1, data, prevBlock.Hash, validator.Name)
	newBlock.Hash = newBlock.calculateHash()
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) ChooseValidator(data int) Validator {
	return findNearestValidator(bc.Validators, data)
}

func main() {
	bc := NewBlockchain()

	for i := 0; i < 110_000_000; i += 10_000_000 {
		validator := bc.ChooseValidator(i)
		bc.AddBlock(i, validator)
		fmt.Printf("Block %s mined by validator %s\n", i+1, validator.Name)
	}

	for _, block := range bc.Blocks {
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Printf("Data: %d\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Validator: %s\n", block.Validator)
	}
}
