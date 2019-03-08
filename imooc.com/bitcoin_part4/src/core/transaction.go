package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
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
	Txid      []byte
	Vout      int
	ScriptSig string //输入签名
}

//TXOutput represents a transaction output
type TXOutput struct {
	Value        int
	ScriptPubKey string //输出签名
}

//NewCoinbaseTX creates a new coinbase transaction
func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	txin := TXInput{[]byte{}, -1, data}                        //发币奖励交易的交易输入是空的
	txout := TXOutput{subsidy, to}                             //固定的发币奖励金额
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}} //结构体数组
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
