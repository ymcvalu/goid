goid is a lib to get the goroutine id, and only support go1.10

# How It Work
the tls save the addr of cur g

use the assembly to acquire the g
```assembly
TEXT ·getG(SB),NOSPLIT,$0-8
	MOVQ (TLS), R14
	MOVQ R14, ret+0(FP)
	RET
``` 
and then count the offset of goid in g

```go
gg := (*g)(nil)
offset = unsafe.Offsetof(gg.goid)
```
we can use the offset and the addr of g to get the go id

we can use runtime.Stack to get the cur go id for checkout:
```go
stack :=make([]byte,12)
runtime.Stack(stack,false)
```
see id_test.go for more info
