package core //可以理解为打包为core类

import (
	"encoding/hex"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"os"
)

const dbFile = "blockchain.db" //打开硬盘上的文件名，会被存放到项目文件夹下
const blocksBucket = "blocks"  //因为db里面有许多桶bucket，[]byte(blockBucket)
const genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink of second bailout of banks"

//原理性展示
type Blockchain struct {
	//Blocks []*Block //存储的是Block结构体的指针的数组

	//现在用一个数据库来做这件事了
	tip []byte //创世区块的哈希值，错，这不是创世区块的哈希值，而是最近的区块的哈希值
	//tail []byte //末尾区块
	Db *bolt.DB
}

//BlockchainIterator is used to iterate over blockchain blocks
type BlockchainIterator struct {
	currentHash []byte   //记录当前的索引/哈希值
	Db          *bolt.DB //访问方法都要变了
}

//Iterator...
func (bc *Blockchain) Iterator() *BlockchainIterator {
	var currentHash []byte
	bc.Db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(blocksBucket))
		//c := b.Cursor()
		//currentHash, _ = c.First()
		//currentHash, _ = c.Last()
		currentHash = b.Get([]byte("t")) //获取末尾的区块
		return nil
	})

	bci := &BlockchainIterator{currentHash, bc.Db}

	return bci //区块链迭代器
}

//Next returns next block starting from the tip
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.Db.View(func(tx *bolt.Tx) error {
		fmt.Printf("in i.Db.View(func(tx *bolt.Tx)\n") //test
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)   //这个地方不应该传递这个吧？
		block = DeserializeBlock(encodedBlock) //这个地方报错了

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	i.currentHash = block.PreBlockHash //

	return block
}

//FindUnspentTransactions returns a list of transactions containing unspent outputs 不考虑是否足够
func (bc *Blockchain) FindUnspentTransactions(address string) []Transaction {
	var unspentTxs []Transaction //包含未花出去的交易输出的交易记录Transaction, 放到数组里面
	spentTXOs := make(map[string][]int)
	bci := bc.Iterator() //获取区块链对象迭代器

	for { //类似C++的while(1) {}
		block := bci.Next() //搜索下一个区块

		for _, tx := range block.Transactions { //遍历一个区块的所有交易数组[]Transaction
			txID := hex.EncodeToString(tx.ID) //交易序号，[]byte转换为string

		Outputs:
			for outIdx, out := range tx.Vout { //单个交易记录Transaction，遍历其所有交易输出[]TXOutput
				//Was the output spent?
				if spentTXOs[txID] != nil { //如果没有OutputX有箭头出来
					for _, spentOut := range spentTXOs[txID] {
						if spentOut == outIdx {
							continue Outputs
						}
					}
				}

				//经过上面的步骤，排除掉了有箭头出来的交易输出TXOutput
				if out.CanBeUnlockedWith(address) { //判断是否是我的地址的钱
					unspentTxs = append(unspentTxs, *tx)
				}
			}

			//发币交易只有一个空的交易输入和一个交易输出
			if tx.IsCoinbase() == false { //如果是普通交易/转账交易
				for _, in := range tx.Vin { //只有普通交易/转账交易有有效的交易输入可找
					if in.CanUnlockOutputWith(address) { //判断这个交易输入是否属于我？
						inTxID := hex.EncodeToString(in.Txid) //交易输入的输入序号
						spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout)
					}
				}
			}
		}

		if len(block.PreBlockHash) == 0 { //说明当前block是创世区块
			break
		}
	}

	return unspentTxs
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
			//c := b.Cursor()
			//preBlockHash, _ := c.Last() //为什么输出是6c?
			//preBlockHash, _ := c.First() //这个是创世区块的哈希值
			//还是用遍历的方法更保险，为什么计算出来的Hash还是6c
			var preBlockHash []byte
			//_, genesisBlockHash := c.Seek([]byte("l"))
			//for k, _ := c.Seek(genesisBlockHash); k != nil; k, _ = c.Next() {
			//	preBlockHash = k
			//}
			preBlockHash = b.Get([]byte("t"))
			//fmt.Printf("preBlockHash: %x\n", preBlockHash)
			newBlock := NewBlock(data, preBlockHash) //创建新区块
			//装桶
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			//更新末尾区块的哈希值
			err = b.Put([]byte("l"), newBlock.Hash) //leader,不是leader,应该是last
			if err != nil {
				log.Panic(err)
			}
			//bc.tail = newBlock.Hash
			bc.tip = newBlock.Hash
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
func NewBlockchain(address string) *Blockchain {
	if dbExists() == false { //要先调用CreateBlockchain函数
		fmt.Println("No existing blockchain found. Create one first.")
		os.Exit(1)
	}

	var tip []byte //创世区块的哈希值
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error { //这是什么用法？
		b := tx.Bucket([]byte(blocksBucket))
		tip = b.Get([]byte("l"))

		return nil //没有error
	})

	if err != nil { //如果上面运行出错了
		log.Panic(err)
	}

	bc := Blockchain{tip, db} //创建区块链

	return &bc
}

//CreateBlockchain create a new blockchain DB
func CreateBlockchain(address string) *Blockchain {
	if dbExists() { //这是什么函数
		fmt.Println("Blockchain alread exits.")
		os.Exit(1)
	}

	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil) //不能重复打开
	if err != nil {
		log.Panic(err)
	}

	//对文件数据库进行更新
	err = db.Update(func(tx *bolt.Tx) error {
		cbtx := NewCoinbaseTX(address, genesisCoinbaseData) //创建一个发币奖励交易记录
		genesis := NewGenesisBlock(cbtx)                    //创建新的创世区块，输入交易记录

		b, err := tx.CreateBucket([]byte(blocksBucket)) //创建桶
		if err != nil {
			log.Panic(err)
		}

		err = b.Put(genesis.Hash, genesis.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), genesis.Hash) //last
		if err != nil {
			log.Panic(err)
		}

		tip = genesis.Hash

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	bc := Blockchain{tip, db}

	return &bc
}

//MineBlock mines a new block with the provided transaction
func (bc *Blockchain) MineBlock(transaction []*Transaction) {
	var lastHash []byte

	err := bc.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		//找到离我们最近的区块的哈希值，新区块要和它挂钩
		lastHash = b.Get([]byte("l")) //tips存放的创世区块的哈希值

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock(transaction, lastHash) //创建了一个包含了新交易记录数组的区块

	//把新的区块挂钩到区块链数据结构里面
	err = bc.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), newBlock.Hash) //现在我是老大
		if err != nil {
			log.Panic(err)
		}

		bc.tip = newBlock.Hash

		return nil
	})
}
