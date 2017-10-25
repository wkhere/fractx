
#include "textflag.h"

// func iter(x0, y0 float64) int64
TEXT ·iter(SB),NOSPLIT,$8-24
    MOVLPD  x0+0(FP), X0
    MOVHPD  y0+8(FP), X0
    MOVLPD  x0+0(FP), X4
    MOVLPD  y0+8(FP), X5

    MOVLPD  ·pbound(SB), X7
    MOVLPD  ·two(SB), X6
    MOVQ    ·maxi(SB), DX
    MOVQ    $1, CX

    // regs:
    // X0 - current point on entry; accumulator
    // X1 - helper; 2nd accumulator
    // X2 - backup of current point
    // X4L - backup of x0
    // X5L - backup of y0
    // X6L - const 2
    // X7L - const pbound = 4

loop:
    MOVAPD  X0, X2
    MULPD   X0, X0          // X0L=x*x, X0H=y*y
    MOVHPD  X0, tmp-8(SP)
    MOVLPD  tmp-8(SP), X1   // X1L = y*y
    MOVDDUP X0, X0          // X0H = x*x
    MOVDDUP X1, X1          // X1H = y*y
    LONG $0xC1D00f66
    //^  ADDSUBPD X1, X0 | X0L -= X1L; X0H += X1H
    // now: X0L = x*x-y*y; X0H = x*x+y*y
    MOVHPD  X0, tmp-8(SP)
    MOVLPD  tmp-8(SP), X1   // X1L = x*x+y*y
    CMPSD   X7, X1, 1       // < 4
    MOVMSKPD X1, BX
    TESTL   $1, BX
    JZ      end

    ADDSD   X4, X0          // X0L = x' = x*x-y*y+x0

    MOVAPD  X2, X1          // X1L = x
    MOVHPD  X2, tmp-8(SP)
    MULSD   tmp-8(SP), X1   // X1L *= y
    MULSD   X6, X1          // X1L *= 2
    ADDSD   X5, X1          // X1L += y0
    MOVLPD  X1, tmp-8(SP)
    MOVHPD  tmp-8(SP), X0   // X0H = y'
    INCQ    CX
    CMPQ    CX, DX
    JL      loop

end:
    MOVQ    CX, ret+16(FP)
    RET
