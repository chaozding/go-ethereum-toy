package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
)

func calculateHash(toBeHashed string) string {
	hashInBytes := sha256.Sum256([]byte(toBeHashed)) //字节数组
	hashInStr := hex.EncodeToString(hashInBytes[:]) //字节切片
	log.Printf("%s, %s", toBeHashed, hashInStr)
	return hashInStr
}

func main() {
	calculateHash("test1")
	calculateHash("test1")
	calculateHash("test2")
}