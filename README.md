# go-libs

Shared internal helpers used across go-* libraries and application modules.

## Install

```bash
go get github.com/andywang2011/go-libs@v0.1.0
```

## Quick start

```go
digest := hash.Md5("hello")
fmt.Println(digest)
```

## Examples

- `cryptox/aesx` — AES ECB/CBC helpers
- `cryptox/rsax` — RSA PKCS#1 / PKCS#8
- `cryptox/ades` — DES, 3DES, and AES block modes
- `cryptox/pgp` — OpenPGP encrypt/decrypt
- `random` — math and crypto random helpers
- `types` — generic `Set` and slice/map utilities
- `timex` — JSON/database time type
- `jsonmap` — GORM JSON map column type
- `reflectx` — struct-to-map reflection
- `hash` — MD5 helper
- `encrypt` — AES-CBC encrypt service (DI-friendly)

```bash
cd go-libs/examples/md5 && go run .
```

## License

MIT — see [LICENSE](LICENSE).
