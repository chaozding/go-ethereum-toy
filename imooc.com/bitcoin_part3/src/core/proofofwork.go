package core

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64 //这个是最大循环次数，超过了就没必要再工作了
)

const targetBits = 20 //这个是工作量目标，目标为20位，比特币系统规定好的
//为什么20时，是5个0？

//ProofOfWork represents a proof-of-work
type ProofOfWork struct {
	block *Block //通过指针调用某个区块，根据目标，计算哈希值
	//target是系统配置的一个工作量目标
	target *big.Int //这个big什么类型，应该是个大Int类型
}

//NewProofOfWork builds and return a ProofOfWork 相当于构造函数
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)                  //初始化为1
	target.Lsh(target, uint(256-targetBits)) //把1左移1位10，左移2位100，左移3位1000
	//为什么是256 -  targetBits 呢？

	pow := &ProofOfWork{b, target} //相当于实例化ProofOfWork为pow，调用的时候会返回给外部pow实例
	return pow
}

//怎么又搞了个准备方法，而不是直接run呢？
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join( //原来这一步是block.SetHash()干的事情
		[][]byte{
			pow.block.PreBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp), //这个是自己写的工具函数
			IntToHex(int64(targetBits)),   //为什么加入这个？反正是固定值，感觉不加也没事？
			IntToHex(int64(nonce)),        //转换位十六进制的字符串
			//把结束条件同时保存起来
		},
		[]byte{}, //本哈希值是空的
	)
	return data //把data归一化为统一的数据结构
}

//不要想，代码懂了，不懂也懂
//Run performs a proof-of-work 类 用到类的成员数据
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce) //嵌套

		hash = sha256.Sum256(data) //输入字节数组，返回的也是字节数组类型
		fmt.Printf("\r%x", hash)   //快速滚动哈希值

		hashInt.SetBytes(hash[:])          //字节数组类型转换为big.Int类型
		if hashInt.Cmp(pow.target) == -1 { //相等？
			break //对比成功
		} else {
			nonce++ //自增到64位整数maxNonce
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:] //这里nonce是条件，hash[:]是结果
}

//Validate validates block's Pow
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)

	hashInt.SetBytes(hash[:])
	isValid := hashInt.Cmp(pow.target) == -1 //努力工作经过证明，则奖励比特币，比特币怎么来的？
	return isValid
}
