package main

import (
	"github.com/boltdb/bolt"
)

// In its essence Blockchain is just a database with certain structure:
// it’s an ordered, back-linked list.
// Which means that blocks are stored in the insertion order
// and that each block is linked to the previous one.
// This structure allows to quickly get the latest block in a chain
// and to (efficiently) get a block by its hash.
type Blockchain struct {
	// blocks []*Block
	tip []byte
	db  *bolt.DB
}

// AddBlock add a new block to Blockchain
// func (bc *Blockchain) AddBlock(data string) {
// 	prevBlock := bc.blocks[len(bc.blocks)-1]
// 	newBlock := NewBlock(data, prevBlock.Hash)
// 	bc.blocks = append(bc.blocks, newBlock)
// }
func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})
	if err != nil {
		panic(err)
	}

	newBlock := NewBlock(data, lastHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		err = b.Put([]byte("l"), newBlock.Hash)
		bc.tip = newBlock.Hash

		return nil
	})
	if err != nil {
		panic(err)
	}
}

// NewGenesisBlock is To add a new block we need an existing block,
// but there’re not blocks in our Blockchain!
// So, in any Blockchain, there must be at least one block,
// and such block, the first in the chain, is called genesis block.
// Let’s implement a method that creates such a block:
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// NewBlockchain creates a Blockchain with the genesis block:
func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			err = b.Put(genesis.Hash, genesis.Serialize())
			err = b.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	bc := Blockchain{tip, db}

	return &bc
}
