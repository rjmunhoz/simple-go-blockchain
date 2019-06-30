package domain

import (
	"encoding/json"
	"log"
	"strings"
	"time"
)

// Block represents a block that wasn't mined yet; thus, it has no hash
type Block struct {
	Sequence     int       `json:"sequence"`
	Data         string    `json:"data"`
	Timestamp    time.Time `json:"timestamp"`
	PreviousHash string    `json:"previousHash"`
}

// MinedBlock represents a block that has already been mined; this, there's a hash and a nonce
type MinedBlock struct {
	Block
	Hash  string `json:"hash"`
	Nonce string `json:"nonce"`
}

// ToJSON converts a block to it's JSON representation
func (block Block) ToJSON() string {
	jsonBlock, err := json.Marshal(block)

	if err != nil {
		log.Fatal(err)
	}

	return string(jsonBlock)
}

// NewBlock creates a new block from a mines block
func NewBlock(previousBlock MinedBlock, data string) Block {
	return Block{Data: data, PreviousHash: previousBlock.Hash, Timestamp: time.Now(), Sequence: previousBlock.Sequence + 1}
}

// CreateGenesis creates a new Genesis block
func CreateGenesis() Block {
	return Block{Sequence: 0, Data: "Gensesis block", Timestamp: time.Now(), PreviousHash: strings.Repeat("0", 64)}
}
