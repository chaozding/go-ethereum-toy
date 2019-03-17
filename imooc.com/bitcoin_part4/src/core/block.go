package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"time"
)

//
type Block struct {
	//定义结构体
	//Index int64 //区块编号
	Timestamp int64 //区块创建时间戳
	//Data         []byte //区块包含的数据，也可以是string，可以放交易描述
	Transactions []*Transaction //一个区块可以有多个交易对象
	PreBlockHash []byte         //前一个区块的哈希值，也可以是string，不过[]byte更简单
	Hash         []byte         //区块自身的哈希值，用于校验区块的数据有效性

	Nonce int //工作量
}

//Serialize serializes the block 强行编码，序列化区块，把区块结构体序列化为[]byte字符数组类型， 便于存储
func (b *Block) Serialize() []byte {
	var result bytes.Buffer            //临时缓存区域
	encoder := gob.NewEncoder(&result) //传入缓存区，便于存放编码后的结构体对象Block

	err := encoder.Encode(b) //编码
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes() //通过自身类库转成了字节数组类型
}

//DeserializeBlock deserializes a block
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d)) //准备好解码了，因为d是编码后的，所以需要读取
	err := decoder.Decode(&block)                 //执行解码并存储到结构体对象里面
	if err != nil {
		log.Panic(err)
	}

	return &block
}

//经过了工作证明的区块
func NewBlock(transactions []*Transaction, preBlockHash []byte) *Block {
	block := &Block{
		time.Now().Unix(),
		transactions,
		preBlockHash,
		[]byte{},
		0}
	//block.SetHash()
	pow := NewProofOfWork(block) //可以理解为类，然后就不用传递数据了
	//可以理解为类的构造函数，用类的方法进行工作量证明操作
	nonce, hash := pow.Run() //就是挖矿了

	block.Hash = hash[:] //满足系统规定工作量条件的哈希值
	block.Nonce = nonce  //可以理解为工作量完成条件

	return block
}

func NewGenesisBlock(coinbase *Transaction) *Block { //返回的是一个区块结构结构体的指针
	//return NewBlock("Genesis Block", []byte{}) //空数据
	return NewBlock([]*Transaction{coinbase}, []byte{}) //只有一个发币交易,用发币交易初始化数组
}

//func (b *Block) SetHash() {
//	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
//	headers := bytes.Join([][]byte{timestamp, b.Data, b.PreBlockHash}, []byte{})
//	hash := sha256.Sum256(headers) //headers其实包含了数据部分，hash类型是字节数组[]byte
//	b.Hash = hash[:]               //转换为字节数组
//}

//计算区块里所有交易的哈希
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID) //func (tx *Transaction) SetID()
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{})) //为什么需要添加一个空的[]byte{}

	return txHash[:]
}
