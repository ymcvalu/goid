package goid

import (
	"testing"
	"sync"
	"runtime"
	"time"
	"strings"
	"strconv"
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
	goid,_ :=strconv.Atoi(strings.Split(string(stack)," ")[1])
	t.Log(goid)
	t.Log(Goid())
}

func TestAll(t *testing.T){
	s :=time.Now()
	for i:=0;i<10000;i++{
		Goid()
	}
	t.Log(time.Now().Sub(s))
	s = time.Now()
	for i:=0;i<10000;i++{
		Gid()
	}
	t.Log(time.Now().Sub(s))

}