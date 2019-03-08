package core

import (
	"bytes"
	"encoding/gob"
	"log"
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

	Nonce int
}

//Serialize serializes the block 强行编码，序列化区块
func (b *Block) Serialize() []byte {
	var result bytes.Buffer //临时缓存区域
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes() //通过自身类库转成了字节数组类型
}

func NewGenesisBlock() *Block { //返回的是一个区块结构结构体的指针
	return NewBlock("Genesis Block", []byte{}) //空数据
}

//经过了工作证明的区块
func NewBlock(data string, preBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), preBlockHash, []byte{}, 0}
	//block.SetHash()
	pow := NewProofOfWork(block) //可以理解为类，然后就不用传递数据了
	//可以理解为类的构造函数
	nonce, hash := pow.Run() //为什么不用pow.Run(block)
	//就是挖矿了

	block.Hash = hash[:] //满足系统规定工作量条件的哈希值
	block.Nonce = nonce  //可以理解为工作量完成条件

	return block
}

//func (b *Block) SetHash() {
//	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
//	headers := bytes.Join([][]byte{timestamp, b.Data, b.PreBlockHash}, []byte{})
//	hash := sha256.Sum256(headers) //headers其实包含了数据部分，hash类型是字节数组[]byte
//	b.Hash = hash[:]               //转换为字节数组
//}

//DeserializeBlock deserializes a block
//为什么要从字节数组解码区块
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d)) //准备好解码了
	err := decoder.Decode(&block)                 //执行解码
	if err != nil {
		log.Panic(err)
	}

	return &block
}
