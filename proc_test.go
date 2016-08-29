package goid

import (
	"runtime"
	"sync"
	"testing"
)

func TestProcID(t *testing.T) {
	r := struct {
		sync.Mutex
		counts map[int32]int
	}{
		counts: make(map[int32]int),
	}

	goMaxProcs := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				pass, id := testProcIDs()
				if !pass {
					t.Errorf("ProcID changed within a running goroutine")
					runtime.Goexit()
				}
				if id < 0 || int(id) >= goMaxProcs {
					t.Errorf("ProcID value %v is out of range [0, GOMAXPROCS=%v)", id, goMaxProcs)
				}
				r.Lock()
				r.counts[id]++
				r.Unlock()
			}
		}()
	}

	wg.Wait()
	if len(r.counts) != goMaxProcs {
		t.Errorf("Only values up to GOMAXPROCS should have counts: %v", r.counts)
	}
}

func BenchmarkProcID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProcID()
	}
}

func testProcIDs() (bool, int32) {
	var lastID int32 = -1
	for i := 0; i < 1000; i++ {
		id := ProcID()
		if lastID != id && lastID != -1 {
			return false, lastID
		}
		lastID = id
	}
	return true, lastID
}
