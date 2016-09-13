package main

import (
	"encoding/base64"
	"fmt"
)

const (
	base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

var coder = base64.NewEncoding(base64Table)

func base64Encode(src []byte) []byte {
	return []byte(coder.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return coder.DecodeString(string(src))
}
func main() {
	st := base64Encode([]byte("你好吗？小子，干嘛不来一起编程了？o"))
	fmt.Print(string(st), "\n")
	stde, _ := base64Decode(st)
	fmt.Println(string(stde))
}
