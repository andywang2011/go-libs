// MD5 helper demo (uses encrypt subpackage for a second scenario).
package main

import (
	"fmt"

	"github.com/andywang2011/go-libs/encrypt"
	"github.com/andywang2011/go-libs/hash"
)

func main() {
	fmt.Println("md5:", hash.Md5("payload"))
	svc := encrypt.NewEncryptServiceFromConfig(encrypt.EncryptConfig{
		Key: "0123456789abcdef0123456789abcdef",
		Iv:  "0123456789abcdef",
	})
	cipher, err := svc.Encrypt("secret")
	if err != nil {
		panic(err)
	}
	fmt.Println("encrypted length:", len(cipher))
}
