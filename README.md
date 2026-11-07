# go-libs

Shared internal helpers used across go-* libraries and application modules.

## Install

```bash
go get github.com/andywang2011/go-libs@v0.1.0
```

## Quick start

```go
digest := hash.SHA256Hex("hello")
fmt.Println(digest)

svc, _ := encrypt.NewEncryptServiceFromConfig(encrypt.EncryptConfig{Key: key32}) // 32-byte AES-256 key
ct, _ := svc.Encrypt("secret")                                                    // base64(nonce||ciphertext||tag)
```

## Examples

- `cryptox/aesx` — AES-GCM (preferred) and AES-CBC with explicit IV
- `cryptox/rsax` — RSA PKCS#1 / PKCS#8
- `cryptox/pgp` — signed OpenPGP encrypt / verified decrypt (ProtonMail go-crypto)
- `random` — CSPRNG helpers (`UUID`, `UUIDWithLen`, `MustToken`, `NumCode`)
- `types` — generic `Set` and slice/map utilities
- `timex` — JSON/database time type
- `jsonmap` — GORM JSON map column type
- `reflectx` — struct-to-map reflection (honours `json:"name,omitempty"`)
- `hash` — SHA-256 and HMAC-SHA256
- `encrypt` — AES-256-GCM encrypt service (DI-friendly)

```bash
cd go-libs/examples/sha256 && go run .
```

## License

MIT — see [LICENSE](LICENSE).
