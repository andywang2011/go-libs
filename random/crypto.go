package random

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
)

// Bytes 随机字节数组
func Bytes(n int) []byte {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return b
}

func UUID() string {
	return hex.EncodeToString(Bytes(16))
}

func UUIDWithLen(n int) string {
	return hex.EncodeToString(Bytes(n / 2))
}

// NumCode 安全的六位数字码
func NumCode() string {
	num, _ := rand.Int(rand.Reader, big.NewInt(1e6))
	return fmt.Sprintf("%06d", num.Int64())
}
