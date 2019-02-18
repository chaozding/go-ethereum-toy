package core

import "log"

//这个区块链得很大吧？
type Blockchain struct {
	Blocks []*Block //结构体，这个是结构体数组吗？
}

//为什么要生成新的区块链，我添加新的数据后不就区块链不就自动更新了吗？
func NewBlockchain() *Blockchain {
	genesisBlock := GenerateGenesisBlock()
	blockchain := Blockchain{} //结构体初始化，这是一个空的
	blockchain.ApendBlock(&genesisBlock)
	return &blockchain //为什么返回引用呢？
}

func (bc *Blockchain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := GenerateNewBlock(*preBlock, data)
	bc.ApendBlock(&newBlock)
}

func (bc *Blockchain) ApendBlock(newBlock *Block) {
	//这个至少要先生成一个新区块吧
	if isValid(*newBlock, *bc.Blocks[len(bc.Blocks)-1]) {
		bc.Blocks = append(bc.Blocks, newBlock) //怎么append的？
	} else {
		log.Fatal("invalid block")
	}
}

func isValid(newBlock Block, oldBlock Block) bool {
	if newBlock.Index-1 != oldBlock.Index {
		return false
	}
	if newBlock.PrevBlockHash != oldBlock.Hash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false // 这一步有必要吗？这个newBlock.Hash本来不就是利用calculateHash(newBlock)计算出来的吗？
	}
	return true
}
