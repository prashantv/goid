// +build go1.5
// +build !go1.6

#include "funcdata.h"
#include "textflag.h"

#define GID_OFSET (23 * 8)
#define M_OFFSET (6 * 8)
#define P_OFFSET (22 * 8)
#define PID_OFFSET (1 * 8)

TEXT ·GoroutineID(SB), NOSPLIT, $0
    MOVQ TLS, CX
	MOVQ 0(CX)(TLS*1), AX
 	MOVQ GID_OFSET(AX), BX
	MOVQ BX, ret+0(FP)
	RET

TEXT ·ProcID(SB), NOSPLIT, $0
    MOVQ TLS, CX
	MOVQ 0(CX)(TLS*1), AX
 	MOVQ M_OFFSET(AX), BX
	MOVQ P_OFFSET(BX), CX
	XORQ BX, BX
	MOVQ PID_OFFSET(CX), BX
	MOVQ BX, ret1+0(FP)
	RET
