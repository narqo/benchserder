// Code generated by command: go run valid_print_asm.go -out valid_print_amd64.s -stubs valid_print_amd64.go. DO NOT EDIT.

#include "textflag.h"

// func validPrintAVX2(p *byte, n uintptr) int
// Requires: AVX, AVX2, SSE, SSE2, SSE4.1
TEXT ·validPrintAVX2(SB), NOSPLIT, $0-24
	MOVQ p+0(FP), AX
	MOVQ n+8(FP), CX
	SHRQ $0x04, CX
	MOVQ $0x0000000000000000, DX

	// Initialize 128 bits registers.
	MOVQ   $0x1919191919191919, BX
	MOVQ   $0x7e7e7e7e7e7e7e7e, BP
	PINSRQ $0x00, BX, X0
	PINSRQ $0x01, BX, X0
	PINSRQ $0x00, BP, X1
	PINSRQ $0x01, BP, X1
	CMPQ   CX, $0x02
	JL     loop16

	// Initialize 256 bits registers.
	VPBROADCASTQ X0, Y2
	VPBROADCASTQ X1, Y3

loop64:
	// Unroll two iterations of the loop operating on 32 bytes chunks.
	CMPQ      CX, $0x04
	JL        loop32
	VMOVUPS   (AX), Y4
	VMOVUPS   32(AX), Y5
	VPCMPGTB  Y2, Y4, Y6
	VPCMPGTB  Y3, Y4, Y4
	VPANDN    Y6, Y4, Y4
	VPCMPGTB  Y2, Y5, Y6
	VPCMPGTB  Y3, Y5, Y5
	VPANDN    Y6, Y5, Y5
	VPAND     Y4, Y5, Y4
	VPMOVMSKB Y4, BX
	XORL      $0xffffffff, BX
	JNE       done
	SUBQ      $0x04, CX
	ADDQ      $0x40, AX
	CMPQ      CX, $0x04
	JGE       loop64

loop32:
	// Consume the next 32 bytes of input.
	CMPQ      CX, $0x02
	JL        loop16
	VMOVUPS   (AX), Y4
	VPCMPGTB  Y2, Y4, Y6
	VPCMPGTB  Y3, Y4, Y4
	VPANDN    Y6, Y4, Y4
	VPMOVMSKB Y4, BX
	XORL      $0xffffffff, BX
	JNE       done
	SUBQ      $0x02, CX
	ADDQ      $0x20, AX

loop16:
	// Consume the next 16 bytes of input.
	CMPQ     CX, $0x00
	JE       valid
	MOVUPS   (AX), X2
	MOVUPS   X2, X3
	PCMPGTB  X0, X2
	PCMPGTB  X1, X3
	PANDN    X2, X3
	PMOVMSKB X2, BX
	XORL     $0x0000ffff, BX
	JNE      done

valid:
	MOVQ $0x00000001, DX

done:
	MOVQ DX, ret+16(FP)
	RET