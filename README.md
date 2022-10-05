# assignment01bca
Assignment-BlockChain-01

Respected Sir, I want to ask you that I did my best to do the code in such a way that would be easily viewable for you. 
Tho i want to say that there are few things that i wish to clarify for you before you mark my assignment!

1. You must call a function called `Initialize_Blockchain()` and populate it `GenesisBlock()`
2. The hashes for are stored in `[]byte` data type so its not `string`
3. For you to add transaction to(or block to blockchain) you will have to call `AddBlock(transaction string)` which 
   in turn will call `CreateBlock(transaction string, nonce int, prev_hash []byte)` <--- what you suggest in assignment description
4. `VerifyChain()` will *ONLY* tell if blockchain has been tempered with or not `reconfigure_blockchain()` *WILL* reconfigure the entire blockchain
5. `GetBlockHash(data string, nonce int, prev_hash []byte) []byte` will calculate the entire hash of a block and then return the hash in byte array format  

