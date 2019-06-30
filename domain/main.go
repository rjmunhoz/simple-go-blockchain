package domain

import (
  "encoding/json"
  "fmt"
  "log"
)

func CreateChain(size int) []MinedBlock {
	var blocks []MinedBlock
	miner := Miner{Difficulty: 4}

	genesis := miner.mine(CreateGenesis())

	blocks = append(blocks, genesis)

	Repeat(func(i int) {
		block := NewBlock(blocks[len(blocks)-1], fmt.Sprintf("Block %d", i))
		blocks = append(blocks, miner.mine(block))
	}, size)

	return blocks
}

func StringifyChain(blockchain []MinedBlock) string {
	jsonBlockchain, err := json.MarshalIndent(blockchain, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	return string(jsonBlockchain)
}

