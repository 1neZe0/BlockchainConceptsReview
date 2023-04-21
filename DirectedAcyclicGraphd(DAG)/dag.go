package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type Transaction struct {
	From   string
	To     string
	Amount int
}

type Block struct {
	Timestamp    string
	Transactions []*Transaction
	PrevHashes   []string
	Hash         string
}

type DAGBlockchain struct {
	Blocks []*Block
}

func NewBlock(transactions []*Transaction, prevHashes []string) *Block {
	block := &Block{
		Timestamp:    time.Now().String(),
		Transactions: transactions,
		PrevHashes:   prevHashes,
	}
	block.Hash = block.calculateHash()
	return block
}

func (b *Block) calculateHash() string {
	record := b.Timestamp
	for _, tx := range b.Transactions {
		record += tx.From + tx.To + strconv.Itoa(tx.Amount)
	}
	for _, prevHash := range b.PrevHashes {
		record += prevHash
	}
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func NewDAGBlockchain() *DAGBlockchain {
	genesisBlock := NewBlock(nil, []string{})
	return &DAGBlockchain{
		Blocks: []*Block{genesisBlock},
	}
}

func (bc *DAGBlockchain) AddBlock(transactions []*Transaction, prevHashes []string) *Block {
	block := NewBlock(transactions, prevHashes)
	bc.Blocks = append(bc.Blocks, block)
	return block
}

func main() {
	bc := NewDAGBlockchain()

	// Add some transactions and blocks
	tx1 := &Transaction{From: "Alice", To: "Bob", Amount: 10}
	block1 := bc.AddBlock([]*Transaction{tx1}, []string{bc.Blocks[0].Hash})

	tx2 := &Transaction{From: "Bob", To: "Charlie", Amount: 5}
	block2 := bc.AddBlock([]*Transaction{tx2}, []string{block1.Hash})

	tx3 := &Transaction{From: "Charlie", To: "Dave", Amount: 3}
	block3 := bc.AddBlock([]*Transaction{tx3}, []string{block1.Hash})

	tx4 := &Transaction{From: "Dave", To: "Eve", Amount: 2}
	block4 := bc.AddBlock([]*Transaction{tx4}, []string{block2.Hash, block3.Hash})
	_ = block4
	// Print the blocks
	fmt.Println("Blocks:")
	for _, block := range bc.Blocks {
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Println("Transactions:")
		for _, tx := range block.Transactions {
			fmt.Printf("From: %s\n", tx.From)
			fmt.Printf("To: %s\n", tx.To)
			fmt.Printf("Amount: %d\n", tx.Amount)
		}
		fmt.Printf("PrevHashes: %v\n", block.PrevHashes)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println()
	}
}
