#include "textflag.h"

// func Get() int64
TEXT ·getG(SB),NOSPLIT,$0-8
	MOVQ (TLS), AX
	MOVQ AX, ret+0(FP)
	RET
