package core //可以理解为打包为core类

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

const dbFile = "blockchain.db" //打开硬盘上的文件名，会被存放到项目文件夹下
const blocksBucket = "blocks"  //因为db里面有许多桶bucket，[]byte(blockBucket)

//原理性展示
type Blockchain struct {
	//Blocks []*Block //存储的是Block结构体的指针的数组

	//现在用一个数据库来做这件事了
	tip []byte //创世区块的哈希值
	Db  *bolt.DB
}

//BlockchainIterator is used to iterate over blockchain blocks
type BlockchainIterator struct {
	currentHash []byte   //记录当前的索引/哈希值
	Db          *bolt.DB //访问方法都要变了
}

//Iterator...
func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.Db}

	return bci //区块链迭代器
}

//Next returns next block starting from the tip
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	i.currentHash = block.PreBlockHash //

	return block
}

//AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data string) {
	//获得最后一个区块，以便获得他的哈希值放到新的区块中
	//preBlock := bc.Blocks[len(bc.Blocks)-1]
	//打开文件
	//fmt.Printf("0\n") //test
	//db, err := bolt.Open(dbFile, 0600, nil) //不可以重复打开数据库对象
	db := bc.Db
	//fmt.Printf("1\n") //test
	//前一个区块的哈希值也是外来数据
	//newBlock := NewBlock(data, preBlock.Hash) //创建新区块
	err := db.Update(func(tx *bolt.Tx) error { //这是什么用法？
		b := tx.Bucket([]byte(blocksBucket)) //查找是否有区块链的桶并取得之
		//fmt.Printf("out\n") //test
		if b != nil { //说明已经存在区块链了
			//fmt.Printf("in\n") //test
			//读取最后一个key
			//Create a cursor for iteration.
			c := b.Cursor()
			//preBlockHash, _ := c.Last() //为什么输出是6c?
			//preBlockHash, _ := c.First() //这个是创世区块的哈希值
			//还是用遍历的方法更保险，为什么计算出来的Hash还是6c
			var preBlockHash []byte
			_, genesisBlockHash := c.Seek([]byte("l"))
			for k, _ := c.Seek(genesisBlockHash); k != nil; k, _ = c.Next() {
				preBlockHash = k
			}
			//fmt.Printf("preBlockHash: %x\n", preBlockHash)
			newBlock := NewBlock(data, preBlockHash) //创建新区块
			//装桶
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
		}
		return nil
	})

	if err != nil { //如果上面运行出错了
		log.Panic(err)
	}

	//bc.Blocks = append(bc.Blocks, newBlock)
}

//func (bc *Blockchain) DeleteBlockchain() {
//	var tx *bolt.Tx
//	tx.db = bc.Db
//	tx.pages = nil
//	err := tx.DeleteBucket([]byte(blocksBucket))
//}

//NewBlockchain create a new Blockchain with genesis Block
//可以理解为区块链数据结构的构造函数吧
func NewBlockchain() *Blockchain {
	//return &Blockchain{[]*Block{NewGenesisBlock()}} //只插入了一个区块
	var tip []byte //干嘛的？
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		//检查是否打开成功
		log.Panic(err)
	}

	//1.调用Update方法进行数据的写入
	err = db.Update(func(tx *bolt.Tx) error { //这是什么用法？
		//2.通过CreateBucket()方法创建BlockBucket(表)，初次使用创建
		b := tx.Bucket([]byte(blocksBucket)) //查找是否有区块链的桶

		//3.通过Put()方法往表里面存储一条数据(key,value)，注意类型必须为[]byte
		if b == nil {
			fmt.Println("No existing blockchain found. Creating a new one...")
			genesis := NewGenesisBlock() //生成创世区块数据，待会放到桶里

			//创建桶，替代了过去的数组
			b, err := tx.CreateBucket([]byte(blocksBucket)) //创建名为blocksBucket的桶
			if err != nil {                                 //创建失败，说明有返回错误类型
				log.Panic(err)
			}

			err = b.Put(genesis.Hash, genesis.Serialize()) //存的是区块
			if err != nil {
				log.Panic(err)
			}

			//存放领头的Hash
			err = b.Put([]byte("l"), genesis.Hash) //leader
			if err != nil {
				log.Panic(err)
			}
			tip = genesis.Hash
		} else {
			//已经存在创世区块走这个分支
			tip = b.Get([]byte("l"))
		}
		return nil //为什么return nil
	})

	if err != nil { //如果上面运行出错了
		log.Panic(err)
	}

	bc := Blockchain{tip, db} //创建区块链

	return &bc
}
