# go-libs

Shared internal helpers across `go-*` libraries and application modules.

## Packages

| Import path | Description |
|-------------|-------------|
| `github.com/andywang2011/go-libs/cryptox/aesx` | AES ECB/CBC helpers |
| `github.com/andywang2011/go-libs/cryptox/rsax` | RSA PKCS1/PKCS8 |
| `github.com/andywang2011/go-libs/cryptox/ades` | DES / 3DES / AES block modes |
| `github.com/andywang2011/go-libs/cryptox/pgp` | OpenPGP encrypt/decrypt |
| `github.com/andywang2011/go-libs/random` | Math and crypto random helpers |
| `github.com/andywang2011/go-libs/types` | Generic helpers (`Set`, slice/map utils) |
| `github.com/andywang2011/go-libs/timex` | JSON/database time type |
| `github.com/andywang2011/go-libs/jsonmap` | GORM JSON map column type |
| `github.com/andywang2011/go-libs/reflectx` | Struct-to-map reflection |
| `github.com/andywang2011/go-libs/hash` | MD5 helper |
| `github.com/andywang2011/go-libs/encrypt` | AES-CBC encrypt service (DI-friendly) |

Do **not** copy `pkg/util` into apps — depend on this module instead. OSS uploads belong in [go-upload](../go-upload).

## Monorepo usage

```go
```go
require github.com/andywang2011/go-libs v0.0.0
replace github.com/andywang2011/go-libs => ../go-libs
```
```
