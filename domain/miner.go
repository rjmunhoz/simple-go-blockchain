package domain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
)

// Miner represents a block miner
type Miner struct {
	Difficulty int
}

func (miner Miner) calculateHash(payload []byte) string {
	hash := sha256.Sum256(payload)
	return hex.EncodeToString(hash[:])
}

func (miner Miner) isHashValid(hash string) bool {
	challenge := strings.Repeat("0", miner.Difficulty)
	return strings.HasPrefix(hash, challenge)
}

func (miner Miner) mine(block Block) MinedBlock {
	hash := ""
	nonce := ""
	payload := block.ToJSON()
	counter := 0

	for hash == "" {
		nonceAttempt := strconv.FormatInt(int64(counter), 16)

		attempt := miner.calculateHash([]byte(payload + nonceAttempt))

		if miner.isHashValid(attempt) {
			hash = attempt
			nonce = nonceAttempt
		}

		counter++
	}

	return MinedBlock{Block: block, Hash: hash, Nonce: nonce}
}
