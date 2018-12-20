package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
  "fmt"

	"github.com/davecgh/go-spew/spew"
)

// Block represents each 'item' in the blockchain
type Block struct {
	Index     int
	Timestamp string
	Data       int
	Hash      string
	PrevHash  string
}

// Blockchain is a series of validated Blocks
var Blockchain []Block

func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + strconv.Itoa(block.Data) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func writeBlock(Data int){
  prevBlock := Blockchain[len(Blockchain)-1]
  newBlock := generateBlock(prevBlock, Data)

  if isBlockValid(newBlock, prevBlock) {
    Blockchain = append(Blockchain, newBlock)
    spew.Dump(Blockchain[len(Blockchain)-1])
  }
}

func generateBlock(oldBlock Block, Data int) Block {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Data = Data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}

// make sure block is valid by checking index, and comparing the hash of the previous block
func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

func main() {
  fmt.Println("Creating Genesis Block: ")
  t := time.Now()
  genesisBlock := Block{}
  genesisBlock = Block{0, t.String(), 0, calculateHash(genesisBlock), ""}
  spew.Dump(genesisBlock)

  Blockchain = append(Blockchain, genesisBlock)
  fmt.Println("\n")
  a:= true

  for a {
    var input int
    fmt.Println("Choose one of the following options: ")
    fmt.Println("1 Dump all Blockchain information")
    fmt.Println("2 Print Last Blcok")
    fmt.Println("3 Add Block")
    fmt.Println("4 Quit")
    fmt.Println("\n")

    fmt.Scan(&input)

    switch input{
      case 1:
        spew.Dump(Blockchain)
      case 2:
        spew.Dump(Blockchain[len(Blockchain)-1])
      case 3:
        var datum int
        fmt.Println("Enter data as int")
        fmt.Scan(&datum)
        writeBlock(datum)
      case 4:
        a = false
    }
  }
  fmt.Println("Goodbye")
}
