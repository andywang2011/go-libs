# go-libs

Reserved module for shared internal helpers across `go-*` libraries.

Most functionality lives in dedicated packages (`go-httplib`, `go-redis`, `go-mongo`, etc.).
Add code here only when it is truly cross-cutting and too small to warrant its own module.
