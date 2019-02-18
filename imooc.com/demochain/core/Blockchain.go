package core

import (
	"fmt"
	"log"
)

//这个区块链得很大吧？
type Blockchain struct {
	Blocks []*Block //结构体，这个是结构体数组吗？因为是通过HASH值构成指针的
}

//为什么要生成新的区块链，我添加新的数据后不就区块链不就自动更新了吗？
func NewBlockchain() *Blockchain { //返回的是链式结构变量的指针
	genesisBlock := GenerateGenesisBlock() //初始化，调用的是包内的函数
	blockchain := Blockchain{}             //初始化了一个空的区块链
	blockchain.ApendBlock(&genesisBlock)   //添加创世区块
	return &blockchain                     //为什么返回引用呢？
}

//向区块链发送数据来添加区块
func (bc *Blockchain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1] //区块
	newBlock := GenerateNewBlock(*preBlock, data)
	bc.ApendBlock(&newBlock)
}

func (bc *Blockchain) ApendBlock(newBlock *Block) {
	if len(bc.Blocks) == 0 { //
		bc.Blocks = append(bc.Blocks, newBlock)
		return //免去了else
	}
	//验证新增区块是否有效，创世区块不需要验证吧？
	if isValid(*newBlock, *bc.Blocks[len(bc.Blocks)-1]) {
		bc.Blocks = append(bc.Blocks, newBlock) //怎么append的？
	} else {
		log.Fatal("invalid block")
	}
}

//用来打印
func (bc *Blockchain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Prev.Hash: %s\n", block.PrevBlockHash)
		fmt.Printf("Curr.Hash: %s\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Println()
	}
}

func isValid(newBlock Block, oldBlock Block) bool {
	if newBlock.Index-1 != oldBlock.Index {
		fmt.Printf("1\n")
		return false
	}
	if newBlock.PrevBlockHash != oldBlock.Hash {
		fmt.Printf("2\n")
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		fmt.Printf("3\n")
		return false // 这一步有必要吗？这个newBlock.Hash本来不就是利用calculateHash(newBlock)计算出来的吗？
	}
	return true
}
