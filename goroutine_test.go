package goid

import (
	"bytes"
	"runtime"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoroutineID(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			assert.Equal(t, goIDStack(), GoroutineID(), "Goroutine ID mismatch")
		}()
	}
}

func BenchmarkGoroutineID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoroutineID()
	}
}

func BenchmarkGoroutineIDStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goIDStack()
	}
}

// goIDStack gets the current goroutine ID using runtime.Stack and parsing the
// goroutine ID from the header.
func goIDStack() int64 {
	const goroutinePrefixLen = len("goroutine ")

	buf := make([]byte, 50)
	runtime.Stack(buf, false)
	buf = buf[goroutinePrefixLen:]
	buf = buf[:bytes.IndexByte(buf, ' ')]
	id, err := strconv.ParseInt(string(buf), 10, 64)
	if err != nil {
		panic(err)
	}
	return id
}
