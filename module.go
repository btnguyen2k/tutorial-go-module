package mylib

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5String calculates MD5 hash of a string
func Md5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
