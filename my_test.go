package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"strconv"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := NewWallet()
	address := wallet.GetAddress()
	// 1NSD8JSKDURYEHHMZamcRfJnZcJvMURYhr
	fmt.Println(string(address))
}

func TestGetBalance(t *testing.T) {
	address := "1KY1ZHxTe4NiXFW3MoCirVHrbxpJNwckF5"
	if !ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := NewBlockchain("3000")
	UTXOSet := UTXOSet{bc}
	defer bc.db.Close()

	balance := 0
	pubKeyHash := Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}

func TestPrintChain(t *testing.T) {
	bc := NewBlockchain("3000")
	defer bc.db.Close()

	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("============ Block %x ============\n", block.Hash)
		fmt.Printf("Height: %d\n", block.Height)
		fmt.Printf("Prev. block: %x\n", block.PrevBlockHash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))
		for _, tx := range block.Transactions {
			fmt.Println(tx)
		}
		fmt.Printf("\n\n")

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

func TestOpenFile(t *testing.T) {
	db, err := bolt.Open("blockchain_3002.db", 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	db.Close()
}
