// SHA-256 / HMAC / AES-GCM demo.
package main

import (
	"fmt"

	"github.com/andywang2011/go-libs/encrypt"
	"github.com/andywang2011/go-libs/hash"
	"github.com/andywang2011/go-libs/random"
)

func main() {
	fmt.Println("sha256:", hash.SHA256Hex("payload"))
	fmt.Println("hmac:", hash.HMACSHA256Hex([]byte("key"), []byte("payload")))
	fmt.Println("request id:", random.MustToken(16))

	svc, err := encrypt.NewEncryptServiceFromConfig(encrypt.EncryptConfig{
		Key: "0123456789abcdef0123456789abcdef", // demo only; load 32 random bytes from a secret store
	})
	if err != nil {
		panic(err)
	}
	cipher, err := svc.Encrypt("secret")
	if err != nil {
		panic(err)
	}
	plain, err := svc.Decrypt(cipher)
	if err != nil {
		panic(err)
	}
	fmt.Println("gcm roundtrip:", plain == "secret")
}
