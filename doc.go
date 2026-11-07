// Package libs is reserved for shared internal helpers used across go-* modules.
//
// Design constraints:
//   - Pure helpers only: no viper, no log.Fatal, no process lifecycle.
//   - No global mutable state; callers pass inputs and own side effects.
//   - Prefer dedicated go-* SDK modules for network clients and IO boundaries.
//   - ConfigDefault: N/A — pure helper packages with no shared Config struct.
//
// Crypto policy: only authenticated or modern primitives are exposed. AES-GCM
// (cryptox/aesx, encrypt), SHA-256/HMAC (hash), CSPRNG-backed tokens (random),
// and signed OpenPGP (cryptox/pgp). ECB, DES, MD5, and math/rand-based token
// helpers were removed.
//
// This module intentionally stays minimal. Prefer dedicated go-* packages for
// production-ready clients (redis, mongo, httplib, etc.) rather than growing
// this umbrella module.
package libs

// Doc returns a short package description for tests and documentation.
func Doc() string {
	return "go-libs reserved module"
}
