package rsax_test

import (
	"crypto"
	"github.com/andywang2011/go-libs/cryptox/rsax"
	"testing"
)

const (
	prikey = "MIIEogIBAAKCAQEAgkoL7abxc+pyyJeMRDMZQcDBQ6WxcKCEZmUeu47LbBKzzrbm/iPI14xUgSc4zGqAfJO+nM+pUbPNx6db06n7Z5WO01VhmllVDNW8tJ53kPUjxPsMwGZWdCl0yKwzB58BTcadL6fRihxYz2MlmrKzkQO3pztZeOc2w3LYcVSLF0cFtDkFFQv57XhOc6Zlvzq0J78wmT7ofHULnVcmP8cYH9uuRH9I8tvZW0qNKMSErJbAVrHWAa0cN4LWdzjqU//uQhi1QQsoJ7ppMM1uEh+KPwPBgcT8ga3CmAGNlqjcGnL5AIEjfMLgoUUJmA5z05NAhQsJ0ZCVTMzrclLvJzcKdwIDAQABAoIBAEyztu4D2CW7ojOBNGx6AKZ7Ccv/YcTCWu/FHCCi4hMPt5c8iiSXgIeU77LRGlkNYfL1L8Q1ZKLazYx6dE7ZxU2n+Ftl6n3plW5m7h8PI6jNCG/QtCMYnFDqjvwMlEYqT6QwBx2lGHDMW2RMJDk8dwYqQBZNlFDfHZhWpiGh+fiLVDCw+8ptQ6NSbrvWeQstSrIrEAyFeoAUTTS5a8YkhBFyj9KgIRccuF/EJMgqg1aMhiNS7A7UKnATadLcQDkYvAf0rQ6pbdhFzMj90DV0i1zk469HFTqB21FKClr5v//tGklGt6oupxrAJUXrbwquiAsS0KquAl3G9qYqhevT+DECgYEA53auETMy4PHNJ1R6R5boE3tTq6TlF7RvT7cshVHq3OUJzoRwF+nvB/XcHLX5U5P6GSamsOuQ8CwAXw9pEMyMzdcg5uTa5qrKNaxIEluPt6v214nuitEcjUbBEFHowqbTCS8kzFQPh7+pxF9iOvDopJu6mumoCHVCYAfWzXLO1rUCgYEAkBnCBZNcE400eZUcqkhJ/UqeVJYlq1JjTv69PLNPJuZSoO7xsXifo/fhaSJkj79DU0DPcf197Uz8qUmKox/oTA9KhvuaYEpriv26qFrRerAjN7tVKP86J8YhAD2z/YUAcCCNUL5PMy8VlXeffU6ZBVj+FCBNoy1x9wS4KqIly/sCgYAs0nOkywhYPRrgOeCwCk4Ir8IlI/plG8DObAdIDfPwML40arsWNmveW8LzIF4WzDoLMsvmyMc5axvrTJ2o7iN6HOW3MKwfLQWZXTJbs7Vh+QAE/yXdZd8kkIYEdggf4ZROW/3BP6n+PKRQbbQk+8/hSHLbWtH6SVLAK+RxxoC9XQKBgBHsEjSNUltgKZGhmRIYsZMeL7kHf6HiOlW92x1sk8CbztUTQt42yWwkTVW9IFK8kAwPuH1N8NaNx7LqYaLOOh8Tm8pUCqxFNUDxM8DyorEmrQqmcv4JWFtQChS0l+lFHHvVoJ3VPZZ5Ga8gVKbKGsAOhHNOXTmZUl5NqPYsepT3AoGALp49ZBbpo6UcbmOifSGESzurGQVc82/CPGNdQ0Jyaiez1D+cDbNasrMUNYkDESSUMsutcLricFjVsTG3+zsISCxpAOwV40pfOVPbrUixP+9tDVXlyaX5ksM0Lq3b1qNC1YahXn04PR+Orl4+HfrSQznAwtwm1JYjD9UeNN84thA="
	pubkey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAgkoL7abxc+pyyJeMRDMZQcDBQ6WxcKCEZmUeu47LbBKzzrbm/iPI14xUgSc4zGqAfJO+nM+pUbPNx6db06n7Z5WO01VhmllVDNW8tJ53kPUjxPsMwGZWdCl0yKwzB58BTcadL6fRihxYz2MlmrKzkQO3pztZeOc2w3LYcVSLF0cFtDkFFQv57XhOc6Zlvzq0J78wmT7ofHULnVcmP8cYH9uuRH9I8tvZW0qNKMSErJbAVrHWAa0cN4LWdzjqU//uQhi1QQsoJ7ppMM1uEh+KPwPBgcT8ga3CmAGNlqjcGnL5AIEjfMLgoUUJmA5z05NAhQsJ0ZCVTMzrclLvJzcKdwIDAQAB"

	formatPrikey = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAgkoL7abxc+pyyJeMRDMZQcDBQ6WxcKCEZmUeu47LbBKzzrbm
/iPI14xUgSc4zGqAfJO+nM+pUbPNx6db06n7Z5WO01VhmllVDNW8tJ53kPUjxPsM
wGZWdCl0yKwzB58BTcadL6fRihxYz2MlmrKzkQO3pztZeOc2w3LYcVSLF0cFtDkF
FQv57XhOc6Zlvzq0J78wmT7ofHULnVcmP8cYH9uuRH9I8tvZW0qNKMSErJbAVrHW
Aa0cN4LWdzjqU//uQhi1QQsoJ7ppMM1uEh+KPwPBgcT8ga3CmAGNlqjcGnL5AIEj
fMLgoUUJmA5z05NAhQsJ0ZCVTMzrclLvJzcKdwIDAQABAoIBAEyztu4D2CW7ojOB
NGx6AKZ7Ccv/YcTCWu/FHCCi4hMPt5c8iiSXgIeU77LRGlkNYfL1L8Q1ZKLazYx6
dE7ZxU2n+Ftl6n3plW5m7h8PI6jNCG/QtCMYnFDqjvwMlEYqT6QwBx2lGHDMW2RM
JDk8dwYqQBZNlFDfHZhWpiGh+fiLVDCw+8ptQ6NSbrvWeQstSrIrEAyFeoAUTTS5
a8YkhBFyj9KgIRccuF/EJMgqg1aMhiNS7A7UKnATadLcQDkYvAf0rQ6pbdhFzMj9
0DV0i1zk469HFTqB21FKClr5v//tGklGt6oupxrAJUXrbwquiAsS0KquAl3G9qYq
hevT+DECgYEA53auETMy4PHNJ1R6R5boE3tTq6TlF7RvT7cshVHq3OUJzoRwF+nv
B/XcHLX5U5P6GSamsOuQ8CwAXw9pEMyMzdcg5uTa5qrKNaxIEluPt6v214nuitEc
jUbBEFHowqbTCS8kzFQPh7+pxF9iOvDopJu6mumoCHVCYAfWzXLO1rUCgYEAkBnC
BZNcE400eZUcqkhJ/UqeVJYlq1JjTv69PLNPJuZSoO7xsXifo/fhaSJkj79DU0DP
cf197Uz8qUmKox/oTA9KhvuaYEpriv26qFrRerAjN7tVKP86J8YhAD2z/YUAcCCN
UL5PMy8VlXeffU6ZBVj+FCBNoy1x9wS4KqIly/sCgYAs0nOkywhYPRrgOeCwCk4I
r8IlI/plG8DObAdIDfPwML40arsWNmveW8LzIF4WzDoLMsvmyMc5axvrTJ2o7iN6
HOW3MKwfLQWZXTJbs7Vh+QAE/yXdZd8kkIYEdggf4ZROW/3BP6n+PKRQbbQk+8/h
SHLbWtH6SVLAK+RxxoC9XQKBgBHsEjSNUltgKZGhmRIYsZMeL7kHf6HiOlW92x1s
k8CbztUTQt42yWwkTVW9IFK8kAwPuH1N8NaNx7LqYaLOOh8Tm8pUCqxFNUDxM8Dy
orEmrQqmcv4JWFtQChS0l+lFHHvVoJ3VPZZ5Ga8gVKbKGsAOhHNOXTmZUl5NqPYs
epT3AoGALp49ZBbpo6UcbmOifSGESzurGQVc82/CPGNdQ0Jyaiez1D+cDbNasrMU
NYkDESSUMsutcLricFjVsTG3+zsISCxpAOwV40pfOVPbrUixP+9tDVXlyaX5ksM0
Lq3b1qNC1YahXn04PR+Orl4+HfrSQznAwtwm1JYjD9UeNN84thA=
-----END RSA PRIVATE KEY-----`
	formatPubkey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAgkoL7abxc+pyyJeMRDMZ
QcDBQ6WxcKCEZmUeu47LbBKzzrbm/iPI14xUgSc4zGqAfJO+nM+pUbPNx6db06n7
Z5WO01VhmllVDNW8tJ53kPUjxPsMwGZWdCl0yKwzB58BTcadL6fRihxYz2MlmrKz
kQO3pztZeOc2w3LYcVSLF0cFtDkFFQv57XhOc6Zlvzq0J78wmT7ofHULnVcmP8cY
H9uuRH9I8tvZW0qNKMSErJbAVrHWAa0cN4LWdzjqU//uQhi1QQsoJ7ppMM1uEh+K
PwPBgcT8ga3CmAGNlqjcGnL5AIEjfMLgoUUJmA5z05NAhQsJ0ZCVTMzrclLvJzcK
dwIDAQAB
-----END PUBLIC KEY-----`
)

func TestParsePKCS1PrivateKey(t *testing.T) {
	_, err := rsax.ParsePKCS1PrivateKey(prikey)
	if err != nil {
		t.Fatal(err)
	}
}

func TestParsePKCS1PrivateKey2(t *testing.T) {
	_, err := rsax.ParsePKCS1PrivateKey(formatPrikey)
	if err != nil {
		t.Fatal(err)
	}
}

func TestParsePKCS8PublicKey(t *testing.T) {
	_, err := rsax.ParsePKCS8PublicKey(pubkey)
	if err != nil {
		t.Fatal(err)
	}
}

func TestParsePKCS8PublicKey2(t *testing.T) {
	_, err := rsax.ParsePKCS8PublicKey(formatPubkey)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEncryptDecryptOAEP(t *testing.T) {
	pub, err := rsax.ParsePKCS8PublicKey(formatPubkey)
	if err != nil {
		t.Fatal(err)
	}
	priv, err := rsax.ParsePKCS1PrivateKey(formatPrikey)
	if err != nil {
		t.Fatal(err)
	}
	data := []byte("foo bar")
	enc, err := rsax.EncryptWithOAEP(data, nil, pub, crypto.SHA256.New())
	if err != nil {
		t.Fatal(err)
	}
	d, err := rsax.DecryptWithOAEP(enc, nil, priv, crypto.SHA256.New())
	if err != nil {
		t.Fatal(err)
	}
	if string(d) != "foo bar" {
		t.Fatal("DecryptWithOAEP fail")
	}
}

func TestSignWithPKCS1v15(t *testing.T) {
	priv, err := rsax.ParsePKCS1PrivateKey(formatPrikey)
	if err != nil {
		t.Fatal(err)
	}
	data := []byte("foo bar")
	_, err = rsax.SignWithPKCS1v15(data, priv, crypto.SHA256)
	if err != nil {
		t.Fatal(err)
	}
}

func TestVerifyPKCS1v15(t *testing.T) {
	pub, err := rsax.ParsePKCS8PublicKey(formatPubkey)
	if err != nil {
		t.Fatal(err)
	}
	sig := "TcGGtSsIwKpaC4FlOLulno4AYVf7drdiNhoZ58dXmzqwTO6wuGzf+1D4g2QmK3A1HFzyXVRo6LgZb9JguKWR7U+ornG4KEbcHjJ90t3CgP52QKOz/Ptcy6O+xFU8qZpMk/DGy6oXJl7Pzc5/SQUfr0r9bFX+EQ2NfR3OLilVtR6v+UBFN2ZnGhNMjRQ8Iju/LD4cp9sPYu/zKXqO30DjXuK6DZhTwfopl9oqtyw7ozEDoyB/cFlbP9zVL4b2aCvg+dX72OjPgv4ljJA3wRiYG/KFbRc9XWd2RHXy+69LLL+40dHFjJDX+qzQ2S2XqBD1/zsNJ6Y/KUoG/ljh47RUNQ=="
	err = rsax.VerifySignWithPKCS1v15("foo bar", sig, pub, crypto.SHA256)
	if err != nil {
		t.Fatal(err)
	}
}
