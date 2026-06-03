package pgp

import (
	"bytes"
	"errors"
	"io"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
)

func EncryptBytes(origBytes []byte, publicKey []byte, blockType string) ([]byte, error) {
	//origBytes 待加密内容
	//publicKey 公钥
	//blockType 加密后标签
	//Read public key
	keyringBuffer := bytes.NewReader(publicKey)
	entityList, err := openpgp.ReadArmoredKeyRing(keyringBuffer)
	if err != nil {
		return nil, err
	}

	//armor encode
	buf := new(bytes.Buffer)
	armorWrite, err := armor.Encode(buf, blockType, nil)
	if err != nil {
		return nil, err
	}

	//encrypt string
	encryptWrite, err := openpgp.Encrypt(armorWrite, entityList, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	_, err = encryptWrite.Write(origBytes)
	if err != nil {
		return nil, err
	}
	err = encryptWrite.Close()
	if err != nil {
		return nil, err
	}

	err = armorWrite.Close()
	if err != nil {
		return nil, err
	}

	encryptBytes, err := io.ReadAll(buf)
	if err != nil {
		return nil, err
	}

	return encryptBytes, nil
}

func DecryptBytes(encBytes, privateKey, passphrase []byte, blockType string) ([]byte, error) {
	//encBytes 待解密内容
	//privateKey 私钥
	//passphrase 私钥口令
	//blockType 密文标签
	//armor string decode
	armorBytes := bytes.NewBuffer(encBytes)
	armorDecoder, err := armor.Decode(armorBytes)
	if err != nil {
		return nil, err
	}
	if armorDecoder.Type != blockType {
		return nil, errors.New("BlockType Invalid")
	}

	// init some vars
	var entity *openpgp.Entity
	var entityList openpgp.EntityList

	//Read private key
	keyringBuffer := bytes.NewReader(privateKey)
	entityList, err = openpgp.ReadArmoredKeyRing(keyringBuffer)
	if err != nil {
		return nil, err
	}
	entity = entityList[0]

	// Get the passphrase and read the private key.
	// Have not touched the encrypted string yet
	entity.PrivateKey.Decrypt(passphrase) //nolint
	for _, subkey := range entity.Subkeys {
		subkey.PrivateKey.Decrypt(passphrase) //nolint
	}

	// Decrypt it with the contents of the private key
	md, err := openpgp.ReadMessage(armorDecoder.Body, entityList, nil, nil)
	if err != nil {
		return nil, err
	}
	decBytes, err := io.ReadAll(md.UnverifiedBody)
	if err != nil {
		return nil, err
	}

	return decBytes, nil
}
