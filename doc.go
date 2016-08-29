// Package goid returns information about where the current code is executing.
// Do not use this package.
package goid

// GoroutineID returns the ID of the currently executing goroutine.
func GoroutineID() int64

// ProcID returns the ID of the specific GOPROC that is executing this goroutine.
// If GOMAXPROCS is set to n, the returned ID may be [0,n).
func ProcID() int32
