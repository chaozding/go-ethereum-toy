package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
)

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)                        //缓存
	err := binary.Write(buff, binary.BigEndian, num) //不是字节吗？为什么用二进制写呢？
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func DataToHash(data []byte) []byte {
	hash := sha256.Sum256(data) //接收字节数组类型的输入数据
	return hash[:]
}
