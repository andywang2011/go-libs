package hash

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 returns the lowercase hex MD5 digest of str.
func Md5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}
