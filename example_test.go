package libs_test

import (
	"fmt"

	"github.com/andywang2011/go-libs/hash"
)

func ExampleMd5() {
	digest := hash.Md5("hello")
	fmt.Println(digest)
	// Output: 5d41402abc4b2a76b9719d911017c592
}
