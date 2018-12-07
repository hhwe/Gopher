package main

// In its essence blockchain is just a database with certain structure:
// it’s an ordered, back-linked list.
// Which means that blocks are stored in the insertion order
// and that each block is linked to the previous one.
// This structure allows to quickly get the latest block in a chain
// and to (efficiently) get a block by its hash.
type BlockChain struct {
	blocks []*Block
}

// AddBlock add a new block to blockchain
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// NewGenesisBlock is To add a new block we need an existing block,
// but there’re not blocks in our blockchain!
// So, in any blockchain, there must be at least one block,
// and such block, the first in the chain, is called genesis block.
// Let’s implement a method that creates such a block:
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// NewBlockChain creates a blockchain with the genesis block:
func NewBlockChain() *BlockChain {
	return &BlockChain{blocks: []*Block{NewGenesisBlock()}}
}
