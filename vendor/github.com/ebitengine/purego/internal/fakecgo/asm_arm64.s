// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"
#include "../abi/abi_arm64.h"

// Called by C code generated by cmd/cgo.
// func crosscall2(fn, a unsafe.Pointer, n int32, ctxt uintptr)
// Saves C callee-saved registers and calls cgocallback with three arguments.
// fn is the PC of a func(a unsafe.Pointer) function.
TEXT crosscall2(SB), NOSPLIT|NOFRAME, $0
/*
 * We still need to save all callee save register as before, and then
 *  push 3 args for fn (R0, R1, R3), skipping R2.
 * Also note that at procedure entry in gc world, 8(RSP) will be the
 *  first arg.
 */
	SUB  $(8*24), RSP
	STP  (R0, R1), (8*1)(RSP)
	MOVD R3, (8*3)(RSP)

	SAVE_R19_TO_R28(8*4)
	SAVE_F8_TO_F15(8*14)
	STP (R29, R30), (8*22)(RSP)

	// Initialize Go ABI environment
	BL runtime·load_g(SB)
	BL runtime·cgocallback(SB)

	RESTORE_R19_TO_R28(8*4)
	RESTORE_F8_TO_F15(8*14)
	LDP (8*22)(RSP), (R29, R30)

	ADD $(8*24), RSP
	RET
