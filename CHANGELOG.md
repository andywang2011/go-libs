# Changelog

## [Unreleased]

## [0.2.0] - 2026-11-07

### Breaking

- Removed `cryptox/aesx` ECB (`ECBEncrypt`, `ECBDecrypt`) and CBC with key-as-IV (`CBCEncrypt`, `CBCDecrypt`)
- Removed the `cryptox/ades` package (DES, 3DES, AES block modes)
- `aesx.PKCS7UnPadding` now returns `([]byte, error)` and rejects malformed padding; `CBC*WithIV` validate IV and ciphertext length
- `encrypt`: dropped `Iv`; `Key` must be 32 bytes; service now uses AES-256-GCM only; `NewEncryptServiceFromConfig` returns `(*EncryptService, error)`
- `random`: removed `Strings`, `Chars`, `Uppers`, `Lowers`, `Numbers`, `RangeNum` (math/rand); `UUID`, `UUIDWithLen`, `NumCode` now return `(string, error)`; `UUIDWithLen(n)` returns exactly `n` chars; added `MustToken(n)`
- `hash`: removed `Md5`; added `SHA256Hex`, `HMACSHA256`, `HMACSHA256Hex`
- `cryptox/pgp`: migrated to `github.com/ProtonMail/go-crypto/openpgp`; `EncryptBytes` signs; `DecryptBytes` requires a valid signature
- `reflectx.Struct2Map` now returns `(map[string]any, error)`

### Security

- Added `aesx.GCMEncrypt` / `GCMDecrypt` (random nonce, `nonce||ciphertext||tag`) and base64 string helpers
- PGP decryption rejects unsigned or unknown-signer messages
- `Shuffle` uses `math/rand/v2` and is documented as non-cryptographic

### Changed

- `examples/md5` renamed to `examples/sha256`

## [0.1.0] - 2026-11-07

### Added

- Subpackages: `cryptox/*`, `random`, `types`, `timex`, `jsonmap`, `reflectx`, `hash`, `encrypt`
- English README, examples, and package documentation
