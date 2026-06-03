package encrypt_test

import (
	"encoding/base64"
	"testing"

	"github.com/andywang2011/go-libs/encrypt"
)

func TestEncryptService_EncryptDecrypt(t *testing.T) {
	key := base64.StdEncoding.EncodeToString([]byte("1234567890123456"))
	iv := base64.StdEncoding.EncodeToString([]byte("1234567890123456"))
	svc := encrypt.NewEncryptServiceFromConfig(encrypt.EncryptConfig{Iv: iv, Key: key})
	enc, err := svc.Encrypt("hello")
	if err != nil {
		t.Fatal(err)
	}
	if enc == "" {
		t.Fatal("expected non-empty ciphertext")
	}
	plain, err := svc.Decrypt(enc)
	if err != nil {
		t.Fatal(err)
	}
	if plain != "hello" {
		t.Fatalf("Decrypt: got %q", plain)
	}
}

func TestEncryptService_EmptyPlaintext(t *testing.T) {
	svc := encrypt.NewEncryptServiceFromConfig(encrypt.EncryptConfig{Iv: "iv", Key: "key"})
	enc, err := svc.Encrypt("")
	if err != nil || enc != "" {
		t.Fatalf("Encrypt empty: enc=%q err=%v", enc, err)
	}
	dec, err := svc.Decrypt("")
	if err != nil || dec != "" {
		t.Fatalf("Decrypt empty: dec=%q err=%v", dec, err)
	}
}

func TestNewEncryptService(t *testing.T) {
	base := &encrypt.EncryptService{Iv: "a", Key: "b"}
	got := encrypt.NewEncryptService(base)
	if got.Iv != "a" || got.Key != "b" {
		t.Fatalf("NewEncryptService: %+v", got)
	}
}
