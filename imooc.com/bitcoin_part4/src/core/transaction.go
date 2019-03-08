package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

const subsidy = 10

//Transaction represents a Bitcoin transaction
type Transaction struct {
	//一个交易包含很多笔
	ID   []byte     //交易序号
	Vin  []TXInput  //多个交易输入结构体
	Vout []TXOutput //多个交易输出结构体
}

//TXInput represents a transaction input
type TXInput struct {
	Txid      []byte //交易输入ID是什么意思？
	Vout      int    //交易输入值
	ScriptSig string //输入签名，为什么这里用来存放输入地址？
}

//TXOutput represents a transaction output
type TXOutput struct {
	//为什么交易输出没有ID
	Value        int
	ScriptPubKey string //输出签名，为什么这里用来存放输出地址？
}

//NewCoinbaseTX creates a new coinbase transaction
func NewCoinbaseTX(to, data string) *Transaction { //发币交易
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	txin := TXInput{[]byte{}, -1, data}                        //发币奖励交易的交易输入是空的
	txout := TXOutput{subsidy, to}                             //固定的发币奖励金额
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}} //结构体数组
	tx.SetID()

	return &tx
}

//NewUTXOTransaction creates a new transaction
func NewUTXOTransaction(from, to string, amount int, bc *Blockchain) *Transaction {
	var inputs []TXInput //多个交易输入
	var outputs []TXOutput

	acc, validOutputs := bc.FindSpendableOutputs(from, amount) //validOutputs存放的是交易输出集合

	if acc < amount {
		log.Panic("ERROR: Not enough funds")
	}

	//Build a list of inputs
	for txid, outs := range validOutputs {
		//为什么会有TxID？
		txID, err := hex.DecodeString(txid)
		if err != nil {
			log.Panic(err)
		}

		//没有交易
		for _, out := range outs { //输出作为输入from
			input := TXInput{txID, out, from} //作为新交易的交易输入
			inputs = append(inputs, input)
		}
	}

	//Build a list of outputs
	outputs = append(outputs, TXOutput{amount, to}) //接收地址得到的钱
	if acc > amount {
		outputs = append(outputs, TXOutput{acc - amount, from}) //a change
	}

	tx := Transaction{nil, inputs, outputs} //结合成一个新的交易记录
	tx.SetID()

	return &tx
}

//SetID sets ID of a transaction
func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:] //字节数组类型，比字符数组范围更广
}
