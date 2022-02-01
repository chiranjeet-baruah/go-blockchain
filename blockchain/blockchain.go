package blockchain

import (
	"log"

	"github.com/go-blockchain/utils"
)

// block represents a block in the blockchain
type block struct {
	Index     int    // The index of the block
	Timestamp int64  // The epoch timestamp of the block
	Data      string // The data contained in the block
	PrevHash  string // The hash of the previous block
	Proof     int    // The proof of work
}

// blockchain is a collection of blocks
type blockchain struct {
	Chain []*block // The chain of blocks
}

// NewBlockchain create new instance of blockchain
func NewBlockchain() blockchain {
	blockchain := blockchain{}
	block := block{
		Index:     0,
		Timestamp: utils.GetEpochTime(),
		Data:      "Genesis Block",
		PrevHash:  "",
		Proof:     0,
	}
	blockchain.Chain = append(blockchain.Chain, &block)
	return blockchain
}

// GetLatestBlock returns the latest block in the chain
func (bc *blockchain) GetLatestBlock() *block {
	return bc.Chain[len(bc.Chain)-1]
}

// CreateBlock creates a new block
func (bc *blockchain) CreateBlock(data string) error {
	// Create a new block
	log.Printf("BLOCKCHAIN_ADD_BLOCK: Creating new block with data %v", data)
	return nil
}

// GetBlockchain returns the eltire blockchain
func (bc *blockchain) GetBlockchain() []*block {
	return bc.Chain
}
