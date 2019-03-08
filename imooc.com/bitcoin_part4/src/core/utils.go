package core

import (
	"bytes"
	"encoding/binary"
	"log"
)

//IntToHex convert an int64 to a byte array
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)                        //缓存
	err := binary.Write(buff, binary.BigEndian, num) //不是字节吗？为什么用二进制写呢？
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
