#include "textflag.h"

// func Get() int64
TEXT ·getG(SB),NOSPLIT,$0-8
	MOVQ (TLS), AX
	MOVQ AX, ret+0(FP)
	RET


TEXT ·Goid(SB),NOSPLIT,$0-8
    MOVQ ·offset(SB),AX
    MOVQ (TLS),BX
    ADDQ BX,AX
    MOVQ (AX),BX
    MOVQ BX,ret+0(FP)
    RET
