package aesx

import (
	"testing"
)

var (
	cbcKey         = "bM21oUaFtjCzZKkR"
	cbcOriginData  = "verystart"
	cbcEncryptData = "qeHndO2U0Nz9Q8bdkAiuKA=="
)

func TestCBCEncrypt(t *testing.T) {
	res, err := CBCEncrypt([]byte(cbcKey), []byte(cbcOriginData))
	if err != nil {
		t.Error(err)
	}

	t.Log(res)
	if res != cbcEncryptData {
		t.Error("cbc encrypt wrong")
	}
}

func TestCBCDecrypt(t *testing.T) {
	res, err := CBCDecrypt(cbcKey, cbcEncryptData)
	if err != nil {
		t.Error(err)
	}
	if res != cbcOriginData {
		t.Error("ecb decrypt wrong")
	}
}
