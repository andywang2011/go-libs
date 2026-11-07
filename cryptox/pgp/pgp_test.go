package pgp_test

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/ProtonMail/go-crypto/openpgp"
	"github.com/ProtonMail/go-crypto/openpgp/armor"
	"github.com/andywang2011/go-libs/cryptox/pgp"
)

const blockType = "PGP MESSAGE"

// Fixed armored key material (GnuPG test key from the openpgp test suites).
const testPrivateKey = `-----BEGIN PGP PRIVATE KEY BLOCK-----
Version: GnuPG v1.4.10 (GNU/Linux)

lQHYBE2rFNoBBADFwqWQIW/DSqcB4yCQqnAFTJ27qS5AnB46ccAdw3u4Greeu3Bp
idpoHdjULy7zSKlwR1EA873dO/k/e11Ml3dlAFUinWeejWaK2ugFP6JjiieSsrKn
vWNicdCS4HTWn0X4sjl0ZiAygw6GNhqEQ3cpLeL0g8E9hnYzJKQ0LWJa0QARAQAB
AAP/TB81EIo2VYNmTq0pK1ZXwUpxCrvAAIG3hwKjEzHcbQznsjNvPUihZ+NZQ6+X
0HCfPAdPkGDCLCb6NavcSW+iNnLTrdDnSI6+3BbIONqWWdRDYJhqZCkqmG6zqSfL
IdkJgCw94taUg5BWP/AAeQrhzjChvpMQTVKQL5mnuZbUCeMCAN5qrYMP2S9iKdnk
VANIFj7656ARKt/nf4CBzxcpHTyB8+d2CtPDKCmlJP6vL8t58Jmih+kHJMvC0dzn
gr5f5+sCAOOe5gt9e0am7AvQWhdbHVfJU0TQJx+m2OiCJAqGTB1nvtBLHdJnfdC9
TnXXQ6ZXibqLyBies/xeY2sCKL5qtTMCAKnX9+9d/5yQxRyrQUHt1NYhaXZnJbHx
q4ytu0eWz+5i68IYUSK69jJ1NWPM0T6SkqpB3KCAIv68VFm9PxqG1KmhSrQIVGVz
dCBLZXmIuAQTAQIAIgUCTasU2gIbAwYLCQgHAwIGFQgCCQoLBBYCAwECHgECF4AA
CgkQO9o98PRieSoLhgQAkLEZex02Qt7vGhZzMwuN0R22w3VwyYyjBx+fM3JFETy1
ut4xcLJoJfIaF5ZS38UplgakHG0FQ+b49i8dMij0aZmDqGxrew1m4kBfjXw9B/v+
eIqpODryb6cOSwyQFH0lQkXC040pjq9YqDsO5w0WYNXYKDnzRV0p4H1pweo2VDid
AdgETasU2gEEAN46UPeWRqKHvA99arOxee38fBt2CI08iiWyI8T3J6ivtFGixSqV
bRcPxYO/qLpVe5l84Nb3X71GfVXlc9hyv7CD6tcowL59hg1E/DC5ydI8K8iEpUmK
/UnHdIY5h8/kqgGxkY/T/hgp5fRQgW1ZoZxLajVlMRZ8W4tFtT0DeA+JABEBAAEA
A/0bE1jaaZKj6ndqcw86jd+QtD1SF+Cf21CWRNeLKnUds4FRRvclzTyUMuWPkUeX
TaNNsUOFqBsf6QQ2oHUBBK4VCHffHCW4ZEX2cd6umz7mpHW6XzN4DECEzOVksXtc
lUC1j4UB91DC/RNQqwX1IV2QLSwssVotPMPqhOi0ZLNY7wIA3n7DWKInxYZZ4K+6
rQ+POsz6brEoRHwr8x6XlHenq1Oki855pSa1yXIARoTrSJkBtn5oI+f8AzrnN0BN
oyeQAwIA/7E++3HDi5aweWrViiul9cd3rcsS0dEnksPhvS0ozCJiHsq/6GFmy7J8
QSHZPteedBnZyNp5jR+H7cIfVN3KgwH/Skq4PsuPhDq5TKK6i8Pc1WW8MA6DXTdU
nLkX7RGmMwjC0DBf7KWAlPjFaONAX3a8ndnz//fy1q7u2l9AZwrj1qa1iJ8EGAEC
AAkFAk2rFNoCGwwACgkQO9o98PRieSo2/QP/WTzr4ioINVsvN1akKuekmEMI3LAp
BfHwatufxxP1U+3Si/6YIk7kuPB9Hs+pRqCXzbvPRrI8NHZBmc8qIGthishdCYad
AHcVnXjtxrULkQFGbGvhKURLvS9WnzD/m1K2zzwxzkPTzT9/Yf06O6Mal5AdugPL
VrM0m72/jnpKo04=
=zNCn
-----END PGP PRIVATE KEY BLOCK-----`

func publicKeyOf(t *testing.T, e *openpgp.Entity) []byte {
	t.Helper()
	var buf bytes.Buffer
	w, err := armor.Encode(&buf, openpgp.PublicKeyType, nil)
	if err != nil {
		t.Fatal(err)
	}
	if err := e.Serialize(w); err != nil {
		t.Fatal(err)
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}

func privateKeyOf(t *testing.T, e *openpgp.Entity) []byte {
	t.Helper()
	var buf bytes.Buffer
	w, err := armor.Encode(&buf, openpgp.PrivateKeyType, nil)
	if err != nil {
		t.Fatal(err)
	}
	// Without re-signing so locked (passphrase-protected) keys can be serialized.
	if err := e.SerializePrivateWithoutSigning(w, nil); err != nil {
		t.Fatal(err)
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}

func newEntity(t *testing.T, name string) *openpgp.Entity {
	t.Helper()
	e, err := openpgp.NewEntity(name, "", name+"@example.com", nil)
	if err != nil {
		t.Fatal(err)
	}
	return e
}

func fixedEntity(t *testing.T) *openpgp.Entity {
	t.Helper()
	el, err := openpgp.ReadArmoredKeyRing(strings.NewReader(testPrivateKey))
	if err != nil {
		t.Fatal(err)
	}
	return el[0]
}

func TestEncryptDecryptBytesRoundTripSelfSigned(t *testing.T) {
	e := newEntity(t, "alice")
	pub, priv := publicKeyOf(t, e), privateKeyOf(t, e)
	plain := []byte("hello pgp")

	enc, err := pgp.EncryptBytes(plain, pub, priv, nil, blockType)
	if err != nil {
		t.Fatalf("EncryptBytes: %v", err)
	}
	got, err := pgp.DecryptBytes(enc, priv, nil, nil, blockType)
	if err != nil {
		t.Fatalf("DecryptBytes: %v", err)
	}
	if !bytes.Equal(got, plain) {
		t.Fatalf("plaintext: got %q, want %q", got, plain)
	}
}

func TestEncryptDecryptBytesSeparateSigner(t *testing.T) {
	recipient, sender := newEntity(t, "bob"), newEntity(t, "carol")
	enc, err := pgp.EncryptBytes([]byte("msg"), publicKeyOf(t, recipient), privateKeyOf(t, sender), nil, blockType)
	if err != nil {
		t.Fatal(err)
	}
	got, err := pgp.DecryptBytes(enc, privateKeyOf(t, recipient), nil, publicKeyOf(t, sender), blockType)
	if err != nil || string(got) != "msg" {
		t.Fatalf("got %q err %v", got, err)
	}

	// Without the sender public key the signer is unknown -> must fail.
	if _, err := pgp.DecryptBytes(enc, privateKeyOf(t, recipient), nil, nil, blockType); !errors.Is(err, pgp.ErrUnknownSigner) {
		t.Fatalf("expected ErrUnknownSigner, got %v", err)
	}
	// A different public key must not verify.
	other := newEntity(t, "mallory")
	if _, err := pgp.DecryptBytes(enc, privateKeyOf(t, recipient), nil, publicKeyOf(t, other), blockType); err == nil {
		t.Fatal("expected error for wrong signer key")
	}
}

func TestEncryptDecryptFixedGnuPGKey(t *testing.T) {
	e := fixedEntity(t)
	pub := publicKeyOf(t, e)
	signer := newEntity(t, "dave")
	enc, err := pgp.EncryptBytes([]byte("legacy"), pub, privateKeyOf(t, signer), nil, blockType)
	if err != nil {
		t.Skipf("fixed legacy key not usable for encryption by go-crypto: %v", err)
	}
	got, err := pgp.DecryptBytes(enc, []byte(testPrivateKey), nil, publicKeyOf(t, signer), blockType)
	if err != nil || string(got) != "legacy" {
		t.Fatalf("got %q err %v", got, err)
	}
}

func TestDecryptRejectsUnsignedMessage(t *testing.T) {
	e := newEntity(t, "erin")
	var buf bytes.Buffer
	aw, _ := armor.Encode(&buf, blockType, nil)
	w, err := openpgp.Encrypt(aw, openpgp.EntityList{e}, nil, nil, nil) // no signer
	if err != nil {
		t.Fatal(err)
	}
	_, _ = w.Write([]byte("unsigned"))
	_ = w.Close()
	_ = aw.Close()

	_, err = pgp.DecryptBytes(buf.Bytes(), privateKeyOf(t, e), nil, nil, blockType)
	if !errors.Is(err, pgp.ErrNotSigned) {
		t.Fatalf("expected ErrNotSigned, got %v", err)
	}
}

func TestEncryptDecryptWithPassphrase(t *testing.T) {
	e := newEntity(t, "frank")
	pass := []byte("s3cret")
	if err := e.EncryptPrivateKeys(pass, nil); err != nil {
		t.Fatal(err)
	}
	pub, priv := publicKeyOf(t, e), privateKeyOf(t, e)

	if _, err := pgp.EncryptBytes([]byte("x"), pub, priv, []byte("wrong"), blockType); err == nil {
		t.Fatal("expected error for wrong signer passphrase")
	}
	enc, err := pgp.EncryptBytes([]byte("x"), pub, priv, pass, blockType)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := pgp.DecryptBytes(enc, priv, []byte("wrong"), nil, blockType); err == nil {
		t.Fatal("expected error for wrong passphrase")
	}
	got, err := pgp.DecryptBytes(enc, priv, pass, nil, blockType)
	if err != nil || string(got) != "x" {
		t.Fatalf("got %q err %v", got, err)
	}
}

func TestEncryptBytesInvalidKeys(t *testing.T) {
	e := newEntity(t, "gina")
	pub, priv := publicKeyOf(t, e), privateKeyOf(t, e)
	if _, err := pgp.EncryptBytes([]byte("x"), []byte("not a key"), priv, nil, blockType); err == nil {
		t.Fatal("expected error for invalid public key")
	}
	if _, err := pgp.EncryptBytes([]byte("x"), pub, []byte("not a key"), nil, blockType); err == nil {
		t.Fatal("expected error for invalid signer key")
	}
	// Public key block used as signer has no private material.
	if _, err := pgp.EncryptBytes([]byte("x"), pub, pub, nil, blockType); err == nil {
		t.Fatal("expected error for signer without private key")
	}
}

func TestEmptyKeyring(t *testing.T) {
	e := newEntity(t, "hank")
	pub, priv := publicKeyOf(t, e), privateKeyOf(t, e)
	var empty bytes.Buffer
	w, _ := armor.Encode(&empty, openpgp.PublicKeyType, nil)
	_ = w.Close()

	if _, err := pgp.EncryptBytes([]byte("x"), empty.Bytes(), priv, nil, blockType); err == nil {
		t.Fatal("expected error for empty recipient keyring")
	}
	enc, err := pgp.EncryptBytes([]byte("x"), pub, priv, nil, blockType)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := pgp.DecryptBytes(enc, empty.Bytes(), nil, nil, blockType); err == nil {
		t.Fatal("expected error for empty private keyring")
	}
	if _, err := pgp.DecryptBytes(enc, priv, nil, empty.Bytes(), blockType); err == nil {
		t.Fatal("expected error for empty signer keyring")
	}
}

func TestDecryptBytesInvalidBlockType(t *testing.T) {
	e := newEntity(t, "ivy")
	pub, priv := publicKeyOf(t, e), privateKeyOf(t, e)
	enc, err := pgp.EncryptBytes([]byte("x"), pub, priv, nil, blockType)
	if err != nil {
		t.Fatal(err)
	}
	_, err = pgp.DecryptBytes(enc, priv, nil, nil, "PGP SIGNATURE")
	if !errors.Is(err, pgp.ErrBlockTypeInvalid) || err.Error() != "BlockType Invalid" {
		t.Fatalf("err: %v", err)
	}
}

func TestDecryptBytesInvalidPrivateKey(t *testing.T) {
	e := newEntity(t, "judy")
	pub, priv := publicKeyOf(t, e), privateKeyOf(t, e)
	enc, _ := pgp.EncryptBytes([]byte("x"), pub, priv, nil, blockType)
	if _, err := pgp.DecryptBytes(enc, []byte("bad key"), nil, nil, blockType); err == nil {
		t.Fatal("expected error")
	}
	// Wrong recipient key cannot decrypt.
	other := newEntity(t, "kate")
	if _, err := pgp.DecryptBytes(enc, privateKeyOf(t, other), nil, nil, blockType); err == nil {
		t.Fatal("expected error for wrong private key")
	}
	// Public-only block has no private key.
	if _, err := pgp.DecryptBytes(enc, pub, nil, nil, blockType); err == nil {
		t.Fatal("expected error for public key used as private key")
	}
}

func TestDecryptBytesInvalidCiphertext(t *testing.T) {
	_, err := pgp.DecryptBytes([]byte("not armored"), []byte(testPrivateKey), nil, nil, blockType)
	if err == nil {
		t.Fatal("expected armor decode error")
	}
}
