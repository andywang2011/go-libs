package pgp_test

import (
	"bytes"
	"testing"

	"github.com/andywang2011/go-libs/cryptox/pgp"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
)

const blockType = "PGP MESSAGE"

// Fixed armored key material (GnuPG test key from golang.org/x/crypto/openpgp tests).
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

func testPublicKey(t *testing.T) []byte {
	t.Helper()
	el, err := openpgp.ReadArmoredKeyRing(bytes.NewReader([]byte(testPrivateKey)))
	if err != nil {
		t.Fatal(err)
	}
	var buf bytes.Buffer
	w, err := armor.Encode(&buf, openpgp.PublicKeyType, nil)
	if err != nil {
		t.Fatal(err)
	}
	if err := el[0].Serialize(w); err != nil {
		t.Fatal(err)
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}

func TestEncryptDecryptBytesRoundTrip(t *testing.T) {
	plain := []byte("hello pgp")
	pub := testPublicKey(t)

	enc, err := pgp.EncryptBytes(plain, pub, blockType)
	if err != nil {
		t.Fatalf("EncryptBytes: %v", err)
	}
	if len(enc) == 0 {
		t.Fatal("empty ciphertext")
	}

	got, err := pgp.DecryptBytes(enc, []byte(testPrivateKey), nil, blockType)
	if err != nil {
		t.Fatalf("DecryptBytes: %v", err)
	}
	if !bytes.Equal(got, plain) {
		t.Fatalf("plaintext: got %q, want %q", got, plain)
	}
}

func TestEncryptBytesInvalidPublicKey(t *testing.T) {
	_, err := pgp.EncryptBytes([]byte("x"), []byte("not a key"), blockType)
	if err == nil {
		t.Fatal("expected error for invalid public key")
	}
}

func TestDecryptBytesInvalidBlockType(t *testing.T) {
	pub := testPublicKey(t)
	enc, err := pgp.EncryptBytes([]byte("x"), pub, blockType)
	if err != nil {
		t.Fatal(err)
	}
	_, err = pgp.DecryptBytes(enc, []byte(testPrivateKey), nil, "PGP SIGNATURE")
	if err == nil {
		t.Fatal("expected block type error")
	}
	if err.Error() != "BlockType Invalid" {
		t.Fatalf("err: %v", err)
	}
}

func TestDecryptBytesInvalidPrivateKey(t *testing.T) {
	pub := testPublicKey(t)
	enc, err := pgp.EncryptBytes([]byte("x"), pub, blockType)
	if err != nil {
		t.Fatal(err)
	}
	_, err = pgp.DecryptBytes(enc, []byte("bad key"), nil, blockType)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestDecryptBytesInvalidCiphertext(t *testing.T) {
	_, err := pgp.DecryptBytes([]byte("not armored"), []byte(testPrivateKey), nil, blockType)
	if err == nil {
		t.Fatal("expected armor decode error")
	}
}
