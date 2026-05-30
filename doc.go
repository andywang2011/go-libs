// Package libs is reserved for shared internal helpers used across go-* modules.
//
// This module intentionally stays minimal. Prefer dedicated go-* packages for
// production-ready clients (redis, mongo, httplib, etc.) rather than growing
// this umbrella module.
package libs

// Doc returns a short package description for tests and documentation.
func Doc() string {
	return "go-libs reserved module"
}
