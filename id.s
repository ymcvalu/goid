#include "textflag.h"
#include "funcdata.h"

TEXT ·getG(SB),NOSPLIT,$32-16
    NO_LOCAL_POINTERS
    MOVQ $0, ret_type+0(FP)
    MOVQ $0, ret_data+8(FP)
    GO_RESULTS_INITIALIZED
    MOVQ (TLS),BX
    MOVQ $type·runtime·g(SB),AX
    //for copy the g
    MOVQ AX,  0(SP)
    MOVQ BX,  8(SP)
    CALL runtime·convT2E(SB)
    MOVQ  16(SP),AX
    MOVQ  24(SP),BX
    MOVQ AX,ret_type+0(FP)
    MOVQ BX,ret_data+8(FP)
    RET

TEXT ·Goid(SB),NOSPLIT,$0-8
    NO_LOCAL_POINTERS
    MOVQ ·offset(SB),AX
    MOVQ (TLS),BX
    ADDQ BX,AX
    MOVQ (AX),BX
    MOVQ BX,ret+0(FP)
    RET
