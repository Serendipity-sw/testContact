package main

import (
	"crypto/md5"
	"fmt"
	"encoding/hex"
)

func main() {
	h := md5.New()
	h.Write([]byte("123456"))                          // 需要加密的字符串为 123456
	fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil))) // 输出加密结果
}
