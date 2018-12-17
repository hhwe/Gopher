package main

import (
	"bytes"
	"encoding/gob"
	"time"
)

// In Bitcoint specification Timestamp, PrevBlockHash, and Hash are block headers,
// which form a separate data structure,
// and transactions (Data in our case) is a separate data structure.
// So we’re mixing them here for simplicity
type Block struct {
	// Timestamp is the current timestamp when the block created
	Timestamp int64
	// Stores the hash of the previous block
	PrevBlockHash []byte
	// The hash of current block
	Hash []byte
	// Actual valuable information containing in the block
	Data []byte
	// nonce is required to verify a proof
	Nonce int
}

// // For now, we’ll just take block fields, concatenate them,
// // and calculate a SHA-256 hash on the concatenated combination.
// func (b *Block) SetHash() {
// 	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
// 	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp},
// 		[]byte{})
// 	hash := sha256.Sum256(headers)
// 	b.Hash = hash[:]
// }

// // Creation of a block
// func NewBlock(data string, prevBlockHash []byte) *Block {
// 	block := Block{
// 		Timestamp:     time.Now().Unix(),
// 		PrevBlockHash: prevBlockHash,
// 		Data:          []byte(data),
// 	}
// 	block.SetHash()
// 	return &block
// }

// Serialize the Block. in BoltDB values can be only of []byte type,
// and we want to store Block structs in the DB.
// We’ll use encoding/gob to serialize the structs.
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		panic(err)
	}

	return result.Bytes()
}

// DeserializeBlock is deserializing BoltDB values and return a Block.
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		panic(err)
	}

	return &block
}

// NewBlock create a new block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Nonce:         0,
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}
