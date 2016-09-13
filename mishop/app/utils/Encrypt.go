package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"io"
)

type EncryptUtil struct {
}

func (self EncryptUtil) Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func (self EncryptUtil) Sha1(str string) string {
	h := sha1.New()
	io.WriteString(h, str)
	return hex.EncodeToString(h.Sum(nil))
}

func (self EncryptUtil) Hmac(str string, hashkey string) string {
	key := []byte(hashkey)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(str))
	return hex.EncodeToString(mac.Sum(nil))
}
