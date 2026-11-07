package rsax

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"hash"
	"strings"
)

func ParsePKCS1PrivateKey(key string) (*rsa.PrivateKey, error) {
	var der []byte
	if strings.Contains(key, "-----BEGIN RSA PRIVATE KEY-----") {
		block, _ := pem.Decode([]byte(key))
		if block == nil {
			return nil, errors.New("private key error")
		}
		der = block.Bytes
	} else {
		b, err := base64.StdEncoding.DecodeString(key)
		if err != nil {
			return nil, err
		}
		der = b
	}

	prikey, err := x509.ParsePKCS1PrivateKey(der)
	if err != nil {
		return nil, err
	}
	return prikey, err
}

func ParsePKCS1PublicKey(key string) (*rsa.PublicKey, error) {
	var der []byte
	if strings.Contains(key, "-----BEGIN RSA PUBLIC KEY-----") {
		block, _ := pem.Decode([]byte(key))
		if block == nil {
			return nil, errors.New("public key error")
		}
		der = block.Bytes
	} else {
		b, err := base64.StdEncoding.DecodeString(key)
		if err != nil {
			return nil, err
		}
		der = b
	}

	return x509.ParsePKCS1PublicKey(der)
}

func ParsePKCS8PrivateKey(key string) (*rsa.PrivateKey, error) {
	var der []byte
	if strings.Contains(key, "-----BEGIN PRIVATE KEY-----") {
		block, _ := pem.Decode([]byte(key))
		if block == nil {
			return nil, errors.New("private key error")
		}
		der = block.Bytes
	} else {
		b, err := base64.StdEncoding.DecodeString(key)
		if err != nil {
			return nil, err
		}
		der = b
	}

	prikey, err := x509.ParsePKCS8PrivateKey(der)
	if err != nil {
		return nil, err
	}
	return prikey.(*rsa.PrivateKey), err
}

func ParsePKCS8PublicKey(key string) (*rsa.PublicKey, error) {
	var der []byte
	if strings.Contains(key, "-----BEGIN PUBLIC KEY-----") {
		block, _ := pem.Decode([]byte(key))
		if block == nil {
			return nil, errors.New("public key error")
		}
		der = block.Bytes
	} else {
		b, err := base64.StdEncoding.DecodeString(key)
		if err != nil {
			return nil, err
		}
		der = b
	}

	pub, err := x509.ParsePKIXPublicKey(der)
	if err != nil {
		return nil, err
	}

	pubkey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("public key error")
	}
	return pubkey, err
}

func SignWithPKCS1v15(data []byte, prikey *rsa.PrivateKey, hash crypto.Hash) (string, error) {
	h := hash.New()
	_, err := h.Write(data)
	if err != nil {
		return "", err
	}
	hashed := h.Sum(nil)
	sig, err := rsa.SignPKCS1v15(rand.Reader, prikey, hash, hashed)
	if err != nil {
		return "", err
	}

	s := base64.StdEncoding.EncodeToString(sig)
	return s, nil
}

func VerifySignWithPKCS1v15(data, sign string, pubkey *rsa.PublicKey, hash crypto.Hash) error {
	sig, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}

	h := hash.New()
	if _, err := h.Write([]byte(data)); err != nil {
		return err
	}
	hashed := h.Sum(nil)
	return rsa.VerifyPKCS1v15(pubkey, hash, hashed, sig)
}

func SignWithPSS(data []byte, prikey *rsa.PrivateKey, hash crypto.Hash) (string, error) {
	h := hash.New()
	h.Write(data)
	sign, err := rsa.SignPSS(rand.Reader, prikey, hash, h.Sum(nil), nil)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sign), nil
}

func VerifySignWithPSS(data, sign string, pubkey *rsa.PublicKey, hash crypto.Hash) error {
	sig, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}

	h := hash.New()
	h.Write([]byte(data))
	return rsa.VerifyPSS(pubkey, hash, h.Sum(nil), sig, nil)
}

func EncryptWithOAEP(src, label []byte, key *rsa.PublicKey, h hash.Hash) (string, error) {
	chunk := key.Size() - 2*h.Size() - 2
	if chunk < 0 {
		return "", errors.New("hash size is too long for this key size")
	}
	end := len(src)
	buf := bytes.NewBuffer(nil)
	l := 0
	for l < end {
		r := l + chunk
		if r > end {
			r = end
		}
		b, err := rsa.EncryptOAEP(h, rand.Reader, key, src[l:r], label)
		if err != nil {
			return "", err
		}
		buf.Write(b)
		l = r
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func DecryptWithOAEP(s string, label []byte, key *rsa.PrivateKey, h hash.Hash) ([]byte, error) {
	src, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	chunk := key.Size()
	end := len(src)
	buf := bytes.NewBuffer(nil)
	l := 0
	for l < end {
		r := l + chunk
		if r > end {
			r = end
		}
		b, err := rsa.DecryptOAEP(h, rand.Reader, key, src[l:r], label)
		if err != nil {
			return nil, err
		}
		buf.Write(b)
		l = r
	}
	return buf.Bytes(), nil
}
