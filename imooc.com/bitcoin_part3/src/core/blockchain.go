package core

//原理性展示
type Blockchain struct {
	Blocks []*Block
}

//AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	//前一个区块的哈希值也是外来数据
	newBlock := NewBlock(data, preBlock.Hash) //创建新区块
	bc.Blocks = append(bc.Blocks, newBlock)
}

//NewBlockchain create a new Blockchain with genesis Block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
