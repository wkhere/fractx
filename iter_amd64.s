// iter(x0, y0 float64, maxi uint) uint
TEXT ·iter(SB),$0-32
    XORPD   X0, X0
    XORPD   X1, X1
    MOVLPD  x0+0(FP), X0
    MOVLPD  y0+8(FP), X1
    MOVUPD  X0, X4
    MOVUPD  X1, X5

    XORPD   X7, X7
    MOVLPD  pbound<>(SB), X7
    MOVQ    maxi+16(FP), CX
    XORQ    AX, AX

    // regs:
    // X0L - x, xx
    // X1L - y
    // X2L - backup x
    // X3L - yy
    // X4L - const x0
    // X5L - const y0
    // X6L - backup acc
    // X7L - const pbound = 4.0
    // AX  - i
    // CX  - const maxi

loop:
    MOVUPD  X0, X2
    MOVUPD  X1, X3

    MULSD   X0, X0      // xx
    MULSD   X3, X3      // yy
    MOVUPD  X3, X6
    ADDSD   X0, X6      // xx+yy
    UCOMISD X7, X6      // < 4
    JAE     end

    SUBSD   X3, X0      // xx-yy
    ADDSD   X4, X0      // +x0 = nx

    MULSD   X2, X1      // x*y
    ADDSD   X1, X1      // *2
    ADDSD   X5, X1      // +y0 = ny
    INCQ    AX
    CMPQ    AX, CX
    JL      loop

end:
    MOVQ    AX, ret+24(FP)
    RET


#include "textflag.h"

DATA    pbound<>+0(SB)/8, $4.0
GLOBL   pbound<>(SB),RODATA,$8

// vim: set expandtab:
