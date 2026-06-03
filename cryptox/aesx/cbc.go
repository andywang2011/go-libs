package aesx

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func CBCEncrypt(keyByt, plainByt []byte) (string, error) {
	block, err := aes.NewCipher(keyByt)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	plainByt = PKCS7Padding(plainByt, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, keyByt[:blockSize])
	crypted := make([]byte, len(plainByt))
	blockMode.CryptBlocks(crypted, plainByt)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

func CBCDecrypt(key, cipherText string) (string, error) {
	keyByt := []byte(key)
	cryptedByt, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(keyByt)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, keyByt[:blockSize])
	plainByt := make([]byte, len(cryptedByt))
	blockMode.CryptBlocks(plainByt, cryptedByt)
	b := PKCS7UnPadding(plainByt)
	return string(b), nil
}

func CBCDecryptWithIV(iv, key, cipherText string) (string, error) {
	ivByt, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return "", err
	}
	keyByt, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}
	encryptedByt, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(keyByt)
	if err != nil {
		return "", err
	}
	plainByt := make([]byte, len(encryptedByt))
	cipher.NewCBCDecrypter(block, ivByt).CryptBlocks(plainByt, encryptedByt)

	b := PKCS7UnPadding(plainByt)
	return string(b), nil
}

func CBCEncryptWithIV(iv, key, plainText string) (string, error) {
	ivByt, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return "", err
	}
	keyByt, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyByt)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	plainByt := PKCS7Padding([]byte(plainText), blockSize)
	blockMode := cipher.NewCBCEncrypter(block, ivByt)
	crypted := make([]byte, len(plainByt))
	blockMode.CryptBlocks(crypted, plainByt)
	return base64.StdEncoding.EncodeToString(crypted), nil
}
