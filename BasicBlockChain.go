package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Blockchain struct {
	// array of pointers, pointing to different blocks
	blocks []*Block
}

type Block struct {
	data      string
	nonce     int
	curr_hash []byte
	prev_hash []byte
}

var chain *Blockchain

func GetRandomNonce() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10000)
}

func GetBlockHash(data string, nonce int, prev_hash []byte) []byte {
	str_nonce := strconv.Itoa(nonce)
	str_combined := data + str_nonce
	str_byte := []byte(str_combined)
	slc_byte_arr1, slc_byte_arr2 := str_byte[:], prev_hash[:]
	slc_cmb_byte_arr := make([]byte, 0)
	slc_cmb_byte_arr = append(slc_byte_arr1, slc_byte_arr2...)
	hash := sha256.Sum256(slc_cmb_byte_arr)
	return hash[:]
}

func CreateBlock(transaction string, nonce int, prev_hash []byte) *Block {
	new_block := &Block{"", 0, []byte{}, []byte{}}
	new_block.nonce = nonce
	new_block.prev_hash = prev_hash
	new_block.data = transaction
	new_block.curr_hash = GetBlockHash(transaction, nonce, prev_hash)
	return new_block
}

func AddBlock(data string) {
	prev_block := chain.blocks[len(chain.blocks)-1]
	nonce := GetRandomNonce()
	new_block := CreateBlock(data, nonce, prev_block.curr_hash)
	chain.blocks = append(chain.blocks, new_block)
}

func GenesisBlock() *Block {
	return CreateBlock("Genesis Block", GetRandomNonce(), []byte{})
}

func Initialize_Blockchain() {
	chain = &Blockchain{[]*Block{GenesisBlock()}}
}

func DisplayBlocks() {
	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.prev_hash)
		fmt.Printf("Data in Block: %s\n", block.data)
		fmt.Printf("Nonce: %d\n", block.nonce)
		fmt.Printf("Current Hash : %x\n", block.curr_hash)
		fmt.Println("--------------------------------------------------\n")
	}
}

func verify_blockchain() bool {
	var counter int = 1
	for _, block := range chain.blocks {
		temp_hash := GetBlockHash(block.data, block.nonce, block.prev_hash)
		result := bytes.Compare(temp_hash, block.curr_hash)
		if result == 0 {
			// fmt.Printf("Old Hash %x\nCurrent Hash %x\n", block.curr_hash, temp_hash)
			fmt.Printf("Block#%d with nonce(%d), verified :)\n", counter, block.nonce)
			counter++
		} else {
			// fmt.Printf("Old Hash %x\nCurrent Hash %x\n", block.curr_hash, temp_hash)
			fmt.Printf("Block#%d with nonce(%d), has been tempered with :(\n", counter, block.nonce)
			return false
		}
		fmt.Println("--------------------------------------------------")
	}
	return true
}

func reconfigure_blockchain() {
	for i := 1; i < len(chain.blocks); i++ {
		chain.blocks[i].curr_hash = GetBlockHash(chain.blocks[i].data,
			chain.blocks[i].nonce, chain.blocks[i].prev_hash)
		chain.blocks[i].prev_hash = chain.blocks[i-1].curr_hash
	}
}

func ChangeBlock() {
	fmt.Print("Enter your Block Number: ")
	var block_position int
	fmt.Scanln(&block_position)
	fmt.Print("Enter your data: ")
	var input string
	fmt.Scanln(&input)
	chain.blocks[block_position].data = input
}

func main() {
	Initialize_Blockchain()

	for true {
		fmt.Println("Choose From Following: ⚆ _ ⚆")
		fmt.Println("> Add Block to Blockchain")
		fmt.Println("> Verify Blockchain")
		fmt.Println("> Display Blockchain")
		fmt.Println("> Update Block")
		fmt.Println("> Exit")
		fmt.Print("Enter your choice: ")
		var input int
		fmt.Scanln(&input)
		switch input {
		case 1:
			fmt.Print("Enter your Transaction: ")
			var data string
			fmt.Scanln(&data)
			AddBlock(data)
			fmt.Println("Block Added To Chain Successfully (づ｡◕‿‿◕｡)づ")
		case 2:
			if verify_blockchain() {
				fmt.Println("$ Blockchain is perfectly fine uwu\n")
			} else {
				fmt.Println("somthing went wrong v_v")
				fmt.Println("Automatically Reconfiguring Blockchain")
				reconfigure_blockchain()
				fmt.Println("Blockchain Reconfingured Successfully (づ￣ ³￣)づ")
			}
		case 3:
			fmt.Println("> Your Blockchain is:\n")
			DisplayBlocks()
		case 4:
			ChangeBlock()
		case 5:
			fmt.Println("Exiting Program (ಥ﹏ಥ)")
			return
		default:
			fmt.Println("Wrong input! Try again :)")
		}
	}

}
