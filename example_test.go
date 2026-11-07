package libs_test

import (
	"fmt"

	"github.com/andywang2011/go-libs/hash"
)

func ExampleSHA256Hex() {
	fmt.Println(hash.SHA256Hex("hello"))
	// Output: 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
}
