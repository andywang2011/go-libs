package encrypt_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/andywang2011/go-libs/encrypt"
)

const testKey = "0123456789abcdef0123456789abcdef"

func newSvc(t *testing.T) *encrypt.EncryptService {
	t.Helper()
	svc, err := encrypt.NewEncryptServiceFromConfig(encrypt.EncryptConfig{Key: testKey})
	if err != nil {
		t.Fatal(err)
	}
	return svc
}

func TestEncryptService_EncryptDecrypt(t *testing.T) {
	svc := newSvc(t)
	enc, err := svc.Encrypt("hello")
	if err != nil {
		t.Fatal(err)
	}
	if enc == "" || enc == "hello" {
		t.Fatalf("unexpected ciphertext %q", enc)
	}
	plain, err := svc.Decrypt(enc)
	if err != nil {
		t.Fatal(err)
	}
	if plain != "hello" {
		t.Fatalf("Decrypt: got %q", plain)
	}
}

func TestEncryptService_NonDeterministic(t *testing.T) {
	svc := newSvc(t)
	a, _ := svc.Encrypt("same")
	b, _ := svc.Encrypt("same")
	if a == b {
		t.Fatal("expected random nonce to produce distinct ciphertexts")
	}
}

func TestEncryptService_TamperDetected(t *testing.T) {
	svc := newSvc(t)
	enc, _ := svc.Encrypt("secret")
	// Replace one character in the base64 body.
	b := []byte(enc)
	if b[len(b)/2] == 'A' {
		b[len(b)/2] = 'B'
	} else {
		b[len(b)/2] = 'A'
	}
	if _, err := svc.Decrypt(string(b)); err == nil {
		t.Fatal("expected tamper error")
	}
	if _, err := svc.Decrypt("not-base64!!"); err == nil {
		t.Fatal("expected base64 error")
	}
}

func TestEncryptService_WrongKey(t *testing.T) {
	svc := newSvc(t)
	other, _ := encrypt.NewEncryptServiceFromConfig(encrypt.EncryptConfig{Key: strings.Repeat("k", 32)})
	enc, _ := svc.Encrypt("secret")
	if _, err := other.Decrypt(enc); err == nil {
		t.Fatal("expected error with wrong key")
	}
}

func TestEncryptService_EmptyPlaintext(t *testing.T) {
	svc := newSvc(t)
	enc, err := svc.Encrypt("")
	if err != nil || enc != "" {
		t.Fatalf("Encrypt empty: enc=%q err=%v", enc, err)
	}
	dec, err := svc.Decrypt("")
	if err != nil || dec != "" {
		t.Fatalf("Decrypt empty: dec=%q err=%v", dec, err)
	}
}

func TestNewEncryptServiceFromConfig_KeySize(t *testing.T) {
	for _, k := range []string{"", "short", strings.Repeat("a", 16), strings.Repeat("a", 24), strings.Repeat("a", 33)} {
		if _, err := encrypt.NewEncryptServiceFromConfig(encrypt.EncryptConfig{Key: k}); !errors.Is(err, encrypt.ErrInvalidKeySize) {
			t.Fatalf("key len %d: got %v", len(k), err)
		}
	}
}

func TestEncryptService_InvalidKeyOnZeroValue(t *testing.T) {
	svc := &encrypt.EncryptService{Key: "short"}
	if _, err := svc.Encrypt("x"); !errors.Is(err, encrypt.ErrInvalidKeySize) {
		t.Fatalf("Encrypt: %v", err)
	}
	if _, err := svc.Decrypt("x"); !errors.Is(err, encrypt.ErrInvalidKeySize) {
		t.Fatalf("Decrypt: %v", err)
	}
}

func TestNewEncryptService(t *testing.T) {
	base := &encrypt.EncryptService{Key: testKey}
	got := encrypt.NewEncryptService(base)
	if got.Key != testKey || got == base {
		t.Fatalf("NewEncryptService: %+v", got)
	}
}
