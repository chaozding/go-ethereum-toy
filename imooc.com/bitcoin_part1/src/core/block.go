package core

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

//
type Block struct {
	//定义结构体
	//Index int64 //区块编号
	Timestamp    int64  //区块创建时间戳
	Data         []byte //区块包含的数据，也可以是string，可以非常长
	PreBlockHash []byte //前一个区块的哈希值，也可以是string，不过[]byte更简单
	Hash         []byte //区块自身的哈希值，用于校验区块的数据有效性
}

func NewBlock(data string, preBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), preBlockHash, []byte{}}
	block.SetHash()
	return block
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{timestamp, b.Data, b.PreBlockHash}, []byte{})
	hash := sha256.Sum256(headers) //headers其实包含了数据部分，hash类型是字节数组[]byte
	b.Hash = hash[:]               //转换为字节数组
}
