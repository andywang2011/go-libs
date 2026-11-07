// Package pgp provides signed OpenPGP encryption and verified decryption helpers
// built on github.com/ProtonMail/go-crypto/openpgp.
package pgp

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/ProtonMail/go-crypto/openpgp"
	"github.com/ProtonMail/go-crypto/openpgp/armor"
)

var (
	// ErrBlockTypeInvalid is returned when the armor block type does not match.
	ErrBlockTypeInvalid = errors.New("BlockType Invalid")
	// ErrEmptyKeyring is returned when a supplied key block contains no entities.
	ErrEmptyKeyring = errors.New("pgp: empty keyring")
	// ErrNotSigned is returned when the decrypted message carries no signature.
	ErrNotSigned = errors.New("pgp: message is not signed")
	// ErrUnknownSigner is returned when the signing key is not in the supplied keyrings.
	ErrUnknownSigner = errors.New("pgp: unknown signer")
	// ErrPrivateKeyMissing is returned when no decryptable private key is available.
	ErrPrivateKeyMissing = errors.New("pgp: private key unavailable")
)

func readKeyring(armored []byte) (openpgp.EntityList, error) {
	el, err := openpgp.ReadArmoredKeyRing(bytes.NewReader(armored))
	if err != nil {
		return nil, err
	}
	if len(el) == 0 {
		return nil, ErrEmptyKeyring
	}
	return el, nil
}

// unlock decrypts the entity's private key material with passphrase when it is encrypted.
func unlock(e *openpgp.Entity, passphrase []byte) error {
	if e.PrivateKey == nil {
		return ErrPrivateKeyMissing
	}
	if e.PrivateKey.Encrypted {
		if err := e.PrivateKey.Decrypt(passphrase); err != nil {
			return fmt.Errorf("pgp: decrypt private key: %w", err)
		}
	}
	for _, sub := range e.Subkeys {
		if sub.PrivateKey != nil && sub.PrivateKey.Encrypted {
			if err := sub.PrivateKey.Decrypt(passphrase); err != nil {
				return fmt.Errorf("pgp: decrypt subkey: %w", err)
			}
		}
	}
	return nil
}

// EncryptBytes encrypts origBytes to every entity in the armored recipientPublicKey
// block and signs the message with the first entity in signerPrivateKey.
//
//   - origBytes: plaintext
//   - recipientPublicKey: armored public key(s)
//   - signerPrivateKey: armored private key used to sign (required)
//   - signerPassphrase: passphrase for signerPrivateKey (may be nil if unencrypted)
//   - blockType: armor block type label, e.g. "PGP MESSAGE"
func EncryptBytes(origBytes, recipientPublicKey, signerPrivateKey, signerPassphrase []byte, blockType string) ([]byte, error) {
	recipients, err := readKeyring(recipientPublicKey)
	if err != nil {
		return nil, err
	}
	signers, err := readKeyring(signerPrivateKey)
	if err != nil {
		return nil, err
	}
	signer := signers[0]
	if err := unlock(signer, signerPassphrase); err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	armorWrite, err := armor.Encode(buf, blockType, nil)
	if err != nil {
		return nil, err
	}
	encryptWrite, err := openpgp.Encrypt(armorWrite, recipients, signer, nil, nil)
	if err != nil {
		return nil, err
	}
	if _, err := encryptWrite.Write(origBytes); err != nil {
		return nil, err
	}
	if err := encryptWrite.Close(); err != nil {
		return nil, err
	}
	if err := armorWrite.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// DecryptBytes decrypts armored ciphertext and verifies the signature.
//
//   - encBytes: armored ciphertext
//   - privateKey: armored recipient private key
//   - passphrase: passphrase for privateKey (may be nil if unencrypted)
//   - signerPublicKey: armored sender public key; when empty, signatures are
//     verified against the entities in privateKey (self-signed messages)
//   - blockType: expected armor block type
//
// It returns an error when the message is unsigned, the signature is invalid,
// or the signer is unknown.
func DecryptBytes(encBytes, privateKey, passphrase, signerPublicKey []byte, blockType string) ([]byte, error) {
	armorDecoder, err := armor.Decode(bytes.NewReader(encBytes))
	if err != nil {
		return nil, err
	}
	if armorDecoder.Type != blockType {
		return nil, ErrBlockTypeInvalid
	}

	keyring, err := readKeyring(privateKey)
	if err != nil {
		return nil, err
	}
	for _, e := range keyring {
		if err := unlock(e, passphrase); err != nil {
			return nil, err
		}
	}

	all := append(openpgp.EntityList(nil), keyring...)
	if len(signerPublicKey) > 0 {
		signers, err := readKeyring(signerPublicKey)
		if err != nil {
			return nil, err
		}
		all = append(all, signers...)
	}

	md, err := openpgp.ReadMessage(armorDecoder.Body, all, nil, nil)
	if err != nil {
		return nil, err
	}
	decBytes, err := io.ReadAll(md.UnverifiedBody)
	if err != nil {
		return nil, err
	}
	// SignatureError / IsSigned are only valid after UnverifiedBody is fully consumed.
	if !md.IsSigned {
		return nil, ErrNotSigned
	}
	if md.SignatureError != nil {
		return nil, fmt.Errorf("pgp: signature verification failed: %w", md.SignatureError)
	}
	if md.SignedBy == nil {
		// A signature is present but its issuer is not in the keyring, so it was never verified.
		return nil, ErrUnknownSigner
	}
	return decBytes, nil
}
