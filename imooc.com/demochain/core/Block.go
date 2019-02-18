package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index         int64  //区块编号
	Timestamp     int64  //区块时间戳，创建时间
	PrevBlockHash string //
	Hash          string //当前区块哈希值，不术语当前区块数据

	Data string //区块数据
}

func calculateHash(b Block) string {
	//当前区块哈希值不要包含
	blockData := string(b.Index) + string(b.Timestamp) + string(b.PrevBlockHash) + string(b.Data)
	hashInBytes := sha256.Sum256([]byte(blockData)) //字符串数据类型转换为字节切片
	return hex.EncodeToString(hashInBytes[:])       //又转换为字符串了
}

func GenerateNewBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.PrevBlockHash = preBlock.Hash
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Data = data

	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

func GenerateGenesisBlock() Block {
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	return GenerateNewBlock(preBlock, "Genesis Block")
}
