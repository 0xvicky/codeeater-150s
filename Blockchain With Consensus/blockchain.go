package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	blockNumber int
	prevHash    string
	hash        string
	nonce       int
	timestamp   int
	data        []byte
}

type Blockchain struct {
	chain []Block
}

const difficulty = 2

// STRUCT IMPLEMENTED FUNCTIONS
// function implemented on Blockchain struct, when we work around Blockchain struct this function will also be there
func (vichain *Blockchain) newBlock(data []byte) Block {

	prevBlock := vichain.chain[len(vichain.chain)-1]

	newBlock := Block{
		blockNumber: prevBlock.blockNumber + 1,
		prevHash:    prevBlock.hash,
		nonce:       0,
		timestamp:   int(time.Now().Unix()),
		data:        data,
	}

	newBlock.hash = mineBlock(&newBlock)
	fmt.Println(newBlock.hash)
	return newBlock
}

// MAIN FUNCTIONS
func createGenesisBlock() Block {
	return Block{
		blockNumber: 0,
		prevHash:    "0",
		hash:        "0",
		nonce:       0,
		timestamp:   int(time.Now().Unix()),
		data:        ([]byte("This is genesis block")),
	}
}

func calculateHash(block Block) string {
	record := fmt.Sprintf("%d%s%d%s%d", block.blockNumber, block.prevHash, block.timestamp, string(block.data), block.nonce)
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

func isValidHash(hash string) bool {
	for i := len(hash) - 1; i > len(hash)-difficulty-1; i-- {
		if hash[i] != '0' {
			return false
		}
	}
	return true
}

func mineBlock(block *Block) string {
	for {
		(*block).nonce++
		// fmt.Printf("%d", block.nonce)
		hash := calculateHash(*block)
		fmt.Printf("%s\n", hash)
		if isValidHash(hash) {
			return hash
		}
	}
}

func main() {
	blockchain := Blockchain{chain: []Block{createGenesisBlock()}}

	blockchain.chain = append(blockchain.chain, blockchain.newBlock([]byte("This is first block")))
	blockchain.chain = append(blockchain.chain, blockchain.newBlock([]byte("This is second blocks")))

	for _, block := range blockchain.chain {
		fmt.Println("=======================================")
		fmt.Printf("Block Number:%d\n", block.blockNumber)
		fmt.Printf("Block Previous Hash :%s\n", block.prevHash)
		fmt.Printf("Block Hash :%s\n", block.hash)
		fmt.Printf("Block Nonce :%d\n", block.nonce)
		fmt.Printf("Timestamp: %d\n", block.timestamp)
		fmt.Printf("Data:%s\n", string(block.data))
		fmt.Println("=======================================")
	}

}
