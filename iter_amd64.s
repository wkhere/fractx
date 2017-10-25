
#include "textflag.h"

// func(x0, y0 float64) int64
TEXT ·iterSSE(SB),NOSPLIT,$0-24
    MOVLPD  x0+0(FP), X0
    MOVLPD  y0+8(FP), X1
    MOVAPD  X0, X4
    MOVAPD  X1, X5

    MOVLPD  ·pbound(SB), X7
    MOVQ    ·maxi(SB), CX
    MOVQ    $1, AX

    // regs:
    // X0L - current x, xx
    // X1L - current y, yy
    // X2L -  backup x
    // X3L -  backup y
    // X4L - backup x0
    // X5L - backup y0
    // X6L - backup xx, acc
    // X7L - const pbound = 4

loop:
    MOVAPD  X0, X2
    MOVAPD  X1, X3

    MULSD   X0, X0      // xx
    MULSD   X1, X1      // yy
    MOVAPD  X0, X6
    ADDSD   X1, X6      // xx+yy
    UCOMISD X7, X6      // < 4
    JNB     end

    SUBSD   X1, X0      // xx-yy
    ADDSD   X4, X0      // +x0 = nx

    MOVAPD  X2, X1      // x
    MULSD   X3, X1      // *y
    ADDSD   X1, X1      // *2
    ADDSD   X5, X1      // +y0 = ny
    INCQ    AX
    CMPQ    AX, CX
    JL      loop

end:
    MOVQ    AX, ret+16(FP)
    RET
