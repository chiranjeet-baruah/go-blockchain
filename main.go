package main

import (
	"log"

	"github.com/go-blockchain/blockchain"
)

// VERSION of this service
var VERSION string = "0.0.1"

func main() {
	log.Printf("MAIN_START: Starting version %v", VERSION)

	// Create a new blockchain
	theBlockchain := blockchain.NewBlockchain()
	latestBlock := theBlockchain.GetLatestBlock()
	log.Println("Latest block data:", latestBlock.Data)
	log.Println("Latest block timestamp:", latestBlock.Timestamp)
}
