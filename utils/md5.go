package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

func MakePassword(plainPassword, salt string) string {
	return MD5Encode(plainPassword + salt)
}

func ValidPassword(plainPassword, salt string, password string) bool {
	return MD5Encode(plainPassword+salt) == password
}
