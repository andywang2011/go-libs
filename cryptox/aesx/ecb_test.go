package aesx

import (
	"testing"
)

var (
	key         = "bM21oUaFtjCzZKkR"
	originData  = "verystart"
	encryptData = "0Xwxt5DO4dT4EvhVQCkVuw=="
)

func TestECBEncrypt(t *testing.T) {
	res, err := ECBEncrypt([]byte(key), []byte(originData))
	if err != nil {
		t.Error(err)
	}

	t.Log(res)
	if res != encryptData {
		t.Error("ecb encrypt wrong")
	}
}

func TestECBDecrypt(t *testing.T) {
	res, err := ECBDecrypt([]byte(key), encryptData)
	if err != nil {
		t.Error(err)
	}
	if res != originData {
		t.Error("ecb decrypt wrong")
	}
}
