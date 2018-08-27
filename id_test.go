package goid

import (
	"testing"
	"runtime"
	"strings"
	"strconv"
)

func TestGoid(t *testing.T) {
	stack := make([]byte, 20)
	runtime.Stack(stack, false)
	goid, _ := strconv.Atoi(strings.Split(string(stack), " ")[1])
	t.Log(goid)
	t.Log(Goid())
 }
