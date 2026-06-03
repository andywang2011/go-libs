package ades_test

import (
	"bytes"
	"github.com/andywang2011/go-libs/cryptox/ades"
	"testing"
)

var (
	key8  = []byte(`CiHQ1Ra6`)
	key16 = []byte(`o9kAYsq9mbfqAcnm`)
	key24 = []byte(`AavUaEkMvFomuj1dlsOPH72X`)
	iv8   = []byte(`Yt3JYxVp`)
	iv16  = []byte(`75oUxGPHYFhglSdY`)
)

var text = [][]byte{
	[]byte(`hello world.`),
	[]byte(`{"age":24}`),
	[]byte(`It's yesterday once more.`),
	[]byte(`My heart will go on ~`),
}

func TestNewDesCryptor(t *testing.T) {
	c, err := ades.NewDesCipher(key8, iv8)
	if err != nil {
		t.Fatal(err)
	}
	do(c, t)
}

func TestNewTripleDesCryptor(t *testing.T) {
	c, err := ades.NewTripleDesCipher(key24, iv8)
	if err != nil {
		t.Fatal(err)
	}
	do(c, t)
}

func TestNewAesCryptor(t *testing.T) {
	c, err := ades.NewAesCipher(key16, iv16)
	if err != nil {
		t.Fatal(err)
	}
	do(c, t)
}

func do(c *ades.Cipher, t *testing.T) {
	encrypted := make([][5]string, 0, 4)
	for _, src := range text {
		encrypted = append(encrypted, [5]string{
			c.EncryptECB(src),
			c.EncryptCBC(src),
			c.EncryptCFB(src),
			c.EncryptOFB(src),
			c.EncryptCTR(src),
		})
	}
	for i, s := range encrypted {
		t.Log("src: ", string(text[i]))
		if decrypted, err := c.DecryptECB(s[0]); err != nil {
			t.Error(err)
		} else {
			if bytes.Equal(decrypted, text[i]) {
				t.Logf("\tEncrypt ECB:\t%s", s[0])
			} else {
				t.Errorf("decrypted:%s", decrypted)
			}
		}
		if decrypted, err := c.DecryptCBC(s[1]); err != nil {
			t.Error(err)
		} else {
			if bytes.Equal(decrypted, text[i]) {
				t.Logf("\tEncrypt CBC:\t%s", s[1])
			} else {
				t.Errorf("decrypted:%s", decrypted)
			}
		}
		if decrypted, err := c.DecryptCFB(s[2]); err != nil {
			t.Error(err)
		} else {
			if bytes.Equal(decrypted, text[i]) {
				t.Logf("\tEncrypt CFB:\t%s", s[2])
			} else {
				t.Errorf("decrypted:%s", decrypted)
			}
		}
		if decrypted, err := c.DecryptOFB(s[3]); err != nil {
			t.Error(err)
		} else {
			if bytes.Equal(decrypted, text[i]) {
				t.Logf("\tEncrypt OFB:\t%s", s[3])
			} else {
				t.Errorf("decrypted:%s", decrypted)
			}
		}
		if decrypted, err := c.DecryptCTR(s[4]); err != nil {
			t.Error(err)
		} else {
			if bytes.Equal(decrypted, text[i]) {
				t.Logf("\tEncrypt CTR:\t%s", s[4])
			} else {
				t.Errorf("decrypted:%s", decrypted)
			}
		}
	}
}
