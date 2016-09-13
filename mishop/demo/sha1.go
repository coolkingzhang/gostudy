package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	//sha1
	h := sha1.New()
	io.WriteString(h, "aaaaaa")
	fmt.Printf("%x\n", h.Sum(nil))
	fmt.Printf(hex.EncodeToString(h.Sum(nil)))

	//hmac ,use sha1
	key := []byte("123456")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte("aaaaaa"))
	fmt.Printf("%x\n", mac.Sum(nil))
	fmt.Printf(hex.EncodeToString(mac.Sum(nil)))
}
