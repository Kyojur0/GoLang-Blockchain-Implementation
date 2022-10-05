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

func (block_chain *Blockchain) AddBlock(data string) {
	prev_block := block_chain.blocks[len(block_chain.blocks)-1]
	nonce := GetRandomNonce()
	new_block := CreateBlock(data, nonce, prev_block.curr_hash)
	block_chain.blocks = append(block_chain.blocks, new_block)
}

func GenesisBlock() *Block {
	return CreateBlock("Genesis Block", GetRandomNonce(), []byte{})
}

func Initialize_Blockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}

func (block_chain *Blockchain) DisplayBlockChain() {
	for _, block := range block_chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.prev_hash)
		fmt.Printf("Data in Block: %s\n", block.data)
		fmt.Printf("Nonce: %d\n", block.nonce)
		fmt.Printf("Current Hash : %x\n", block.curr_hash)
		fmt.Println("--------------------------------------------------\n")
	}
}

func (block_chain *Blockchain) verify_blockchain() bool {
	for _, block := range block_chain.blocks {
		temp_hash := GetBlockHash(block.data, block.nonce, block.prev_hash)
		result := bytes.Compare(temp_hash, block.curr_hash)
		if result == 0 {
			fmt.Printf("Old Hash %x\nCurrent Hash %x\n", block.curr_hash, temp_hash)
			fmt.Printf("Block with nonce(%d), verified :)\n", block.nonce)
		} else {
			fmt.Printf("Old Hash %x\nCurrent Hash %x\n", block.curr_hash, temp_hash)
			fmt.Printf("Block with nonce(%d), has been tempered with :(\n", block.nonce)
			return false
		}
		fmt.Println("--------------------------------------------------")
	}
	return true
}

func (block_chain *Blockchain) reconfigure_blockchain() {
	for i := 1; i < len(block_chain.blocks); i++ {
		block_chain.blocks[i].curr_hash = GetBlockHash(block_chain.blocks[i].data,
			block_chain.blocks[i].nonce, block_chain.blocks[i].prev_hash)
		block_chain.blocks[i].prev_hash = block_chain.blocks[i-1].curr_hash
	}
}

func (block_chain *Blockchain) ChangeBlock(block_number int) {
	fmt.Print("Enter your data: ")
	var input string
	fmt.Scanln(&input)
	block_chain.blocks[block_number].data = input
}

func main() {
	chain := Initialize_Blockchain()

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
			chain.AddBlock(data)
			fmt.Println("Block Added To Chain Successfully (づ｡◕‿‿◕｡)づ")
		case 2:
			if chain.verify_blockchain() {
				fmt.Println("Blockchain is perfectly fine uwu")
			} else {
				fmt.Println("somthing went wrong v_v")
				fmt.Println("Automatically Reconfiguring Blockchain")
				chain.reconfigure_blockchain()
				fmt.Println("Blockchain Reconfingured Successfully (づ￣ ³￣)づ")
			}
		case 3:
			fmt.Println("> Your Blockchain is:\n")
			chain.DisplayBlockChain()
		case 4:
			fmt.Print("Enter your Block Number: ")
			var block_position int
			fmt.Scanln(&block_position)
			chain.ChangeBlock(block_position)
		case 5:
			fmt.Println("Exiting Program (ಥ﹏ಥ)")
			return
		default:
			fmt.Println("Wrong input! Try again :)")
		}
	}

}
