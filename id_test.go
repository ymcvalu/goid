package goid

import (
	"testing"
	"sync"
	"runtime"
)

func TestG(t *testing.T){
	gn :=10
	wg :=sync.WaitGroup{}
	wg.Add(gn)
	for i:=0;i<gn;i++{
		go func(){
			stack :=make([]byte,12)
			runtime.Stack(stack,false)
			t.Log(string(stack))
			t.Log(Gid())
			wg.Done()
		}()
	}
	wg.Wait()
}
