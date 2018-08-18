package goid

import (
	"testing"
	"sync"
	"runtime"
)

func TestGid(t *testing.T) {
	gn := 10
	wg := sync.WaitGroup{}
	wg.Add(gn)
	for i := 0; i < gn; i++ {
		go func() {
			stack := make([]byte, 12)
			runtime.Stack(stack, false)
			t.Log(string(stack))
			t.Log(Gid())
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestGoid(t *testing.T) {
	stack := make([]byte, 20)
	runtime.Stack(stack, false)
	t.Log(string(stack))
	t.Log(Goid())
}
