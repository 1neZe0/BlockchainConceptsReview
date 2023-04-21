package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const difficulty = 2

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
	Nonce     int
}

func NewBlock(index int, data string, prevHash string) *Block {
	block := &Block{
		Index:     index,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevHash,
		Nonce:     0,
	}
	block.Hash = block.calculateHash()
	return block
}

func (b *Block) calculateHash() string {
	record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PrevHash + strconv.Itoa(b.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	block := NewBlock(0, "Genesis Block", "")
	return &Blockchain{
		Blocks: []*Block{block},
	}
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(prevBlock.Index+1, data, prevBlock.Hash)
	proof := bc.proofOfWork(newBlock)
	newBlock.Nonce = proof
	newBlock.Hash = newBlock.calculateHash()
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) proofOfWork(block *Block) int {
	var nonce int
	for {
		block.Nonce = nonce
		hash := block.calculateHash()
		if hash[:difficulty] == strings.Repeat("0", difficulty) {
			fmt.Println("Block mined:", hash)
			return nonce
		}
		nonce++
	}
}

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Block 1 Data")
	bc.AddBlock("Block 2 Data")
	bc.AddBlock("Block 3 Data")

	for _, block := range bc.Blocks {
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
	}
}
