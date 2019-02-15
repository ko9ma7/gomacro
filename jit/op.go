/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * op.go
 *
 *  Created on Feb 10, 2019
 *      Author Massimiliano Ghilardi
 */

package jit

import (
	"github.com/cosmos72/gomacro/jit/asm"
)

type Op1 uint8 // unary expression operator
type Op2 uint8 // binary expression operator

const (
	ADD     = Op2(asm.ADD3)
	SUB     = Op2(asm.SUB3)
	MUL     = Op2(asm.MUL3)
	QUO     = Op2(asm.DIV3)
	REM     = Op2(asm.REM3)
	AND     = Op2(asm.AND3)
	OR      = Op2(asm.OR3)
	XOR     = Op2(asm.XOR3)
	SHL     = Op2(asm.SHL3)
	SHR     = Op2(asm.SHR3)
	AND_NOT = Op2(asm.AND_NOT3) // &^
	LAND    = Op2(asm.LAND3)    // &&
	LOR     = Op2(asm.LOR3)     // ||
	/*
		EQL     = Op2(asm.EQL3)
		LSS     = Op2(asm.LSS3)
		GTR     = Op2(asm.GTR3)
		NEQ     = Op2(asm.NEQ3)
		LEQ     = Op2(asm.LEQ3)
		GEQ     = Op2(asm.GEQ3)
	*/
	// CAST = Op1(asm.CAST)
	NEG = Op1(asm.NEG2)
	NOT = Op1(asm.NOT2)
)

func (op Op1) String() string {
	return asm.Op2(op).String()
}

func (op Op2) String() string {
	return asm.Op3(op).String()
}

func (op Op2) IsCommutative() bool {
	return asm.Op3(op).IsCommutative()
}
