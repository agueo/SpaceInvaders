package main

import "fmt"

/*
 * Disassemble8080Op
 * Emits assembly in the form `op src, dest`
 * @param rom - slice of bytes holding rom data
 * @param PC - current PC of program
 * @return int - size of op in bytes
 */
func Disassemble8080Op(rom []byte, PC int) int {
	opsize := 1
	switch op := rom[PC]; op {
	case 0x00:
		fmt.Println("NOP")
	case 0x01:
		// B = byte 3, C = byte 2
		fmt.Println("LDI D16, BC")
		opsize = 3
	case 0x02:
		// (BC) = A
		fmt.Println("STR A (BC)")
	case 0x03:
		// BC = BC + 1
		fmt.Println("INC BC")
	case 0x04:
		// B = B + 1
		fmt.Println("INC B")
	case 0x05:
		// B = B - 1
		fmt.Println("DEC B")
	case 0x06:
		// B = byte 2
		fmt.Println("MVI D8, B")
		opsize = 2
	case 0x07:
		// 	A = A << 1; bit 0 = prev bit 7; CY = prev bit 7
		fmt.Println("RLC")
	case 0x09:
		// HL = HL + BC
		fmt.Println("DAD BC")
	case 0x0a:
		// A = (BC)
		fmt.Println("LD (BC) A")
	case 0x0b:
		// BC = BC - 1
		fmt.Println("DEC BC")
	case 0x0c:
		// C = C + 1
		fmt.Println("INC C")
	case 0x0d:
		// C = C - 1
		fmt.Println("DEC C")
	case 0x0e:
		// C = byte 2
		fmt.Println("MVI D8, C")
		opsize = 2
	case 0x0f:
		// A = A >> 1 bit 7 = prev bit 0
		fmt.Println("RRC")
	case 0x11:
		// D = byte 3, E = byte 2
		fmt.Println("LDI D16, DE")
		opsize = 3
	case 0x12:
		// (DE) = A
		fmt.Println("STR A, (DE)")
	case 0x13:
		// DE = DE + 1
		fmt.Println("INC DE")
	case 0x14:
		// D = D + 1
		fmt.Println("INC D")
	case 0x15:
		// D = D - 1
		fmt.Println("DEC D")
	case 0x16:
		// D = byte 2
		fmt.Println("MVI D8, D")
		opsize = 2
	case 0x17:
		// A = A << 1
		fmt.Println("RAL")
	case 0x19:
		// HL + HL + DE
		fmt.Println("DAD DE")
	case 0x1a:
		// A = (DE)
		fmt.Println("LD A, (DE)")
	case 0x1b:
		// DE = DE - 1
		fmt.Println("DEC DE")
	case 0x1c:
		// E = E + 1
		fmt.Println("INC E")
	case 0x1d:
		// E = E - 1
		fmt.Println("DEC E")
	case 0x1e:
		// E = byte 2
		fmt.Println("MVI D8, E")
		opsize = 2
	case 0x1f:
		// A = A >> 1
		fmt.Println("RAR")
	case 0x20:
		// Special
		fmt.Println("RIM")
	case 0x21:
		// H = byte 3, L = byte 2
		fmt.Println("LDI D16, HL")
		opsize = 3
	case 0x22:
		// (adr) = L, (adr + 1) = H
		fmt.Println("SHLD H, D16")
		opsize = 3
	case 0x23:
		// HL = HL + 1
		fmt.Println("INC HL")
	case 0x24:
		// H = H + 1
		fmt.Println("INC H")
	case 0x25:
		// H = H - 1
		fmt.Println("DEC H")
	case 0x26:
		// L = byte 2
		fmt.Println("MVI L, D8")
		opsize = 2
	case 0x27:
		// H = byte 2
		fmt.Println("DAA")
	case 0x29:
		// HL = HL + HL
		fmt.Println("DAD H")
	case 0x2a:
		// L = (byte 2); H = (byte 3)
		fmt.Println("LHLD adr")
		opsize = 3
	case 0x2b:
		// HL = HL - 1
		fmt.Println("DEC HL")
	case 0x2c:
		// L = L + 1
		fmt.Println("INC L")
	case 0x2d:
		// L = L - 1
		fmt.Println("DEC L")
	case 0x2e:
		// L = byte 2
		fmt.Println("MVI L, D8")
		opsize = 2
	case 0x2f:
		// A = !A
		fmt.Println("CMA")
	case 0x30:
		// SIM - special
		fmt.Println("SIM")
	case 0x31:
		// SP.hi = byte 3, SP.lo = byte 2
		fmt.Println("LDI D16, SP")
		opsize = 3
	case 0x32:
		// (addr) = A
		fmt.Println("STA addr")
		opsize = 3
	case 0x33:
		// SP = SP + 1
		fmt.Println("INC SP")
	case 0x34:
		// (HL) = (HL) + 1
		fmt.Println("INC (HL)")
	case 0x35:
		// (HL) = (HL) - 1
		fmt.Println("DEC (HL)")
	case 0x36:
		// (HL) = byte 2
		fmt.Println("MVI D8 (HL)")
		opsize = 2
	case 0x37:
		// CY = 1
		fmt.Println("STC")
	case 0x39:
		// HL = HL + SP
		fmt.Println("DAD SP")
	case 0x3a:
		// A = (Addr)
		fmt.Println("LDA adr")
		opsize = 3
	case 0x3b:
		// SP = SP - 1
		fmt.Println("DEC SP")
	case 0x3c:
		// A = A + 1
		fmt.Println("INC A")
	case 0x3d:
		// A = A - 1
		fmt.Println("DEC A")
	case 0x3e:
		// A = byte 2
		fmt.Println("MVI D8, A")
		opsize = 2
	case 0x3f:
		// CY != CY
		fmt.Println("CMP C")
	case 0x40:
		// B = B
		fmt.Println("MOV B, B")
	case 0x41:
		// B = C
		fmt.Println("MOV C, B")
	case 0x42:
		// B = D
		fmt.Println("MOV D, B")
	case 0x43:
		// B = H
		fmt.Println("MOV E, B")
	case 0x44:
		// B = B
		fmt.Println("MOV H, B")
	case 0x45:
		// B = B
		fmt.Println("MOV L, B")
	case 0x46:
		// B = B
		fmt.Println("MOV (HL), B")
	case 0x47:
		// B = B
		fmt.Println("MOV A, B")
	case 0x48:
		// B = B
		fmt.Println("MOV B, C")
	case 0x49:
		// B = B
		fmt.Println("MOV C, C")
	case 0x4a:
		// B = B
		fmt.Println("MOV D, C")
	case 0x4b:
		// B = B
		fmt.Println("MOV E, C")
	case 0x4c:
		// B = B
		fmt.Println("MOV H, C")
	case 0x4d:
		// B = B
		fmt.Println("MOV L, C")
	case 0x4e:
		// B = B
		fmt.Println("MOV (HL), C")
	case 0x4f:
		// B = B
		fmt.Println("MOV A, C")
	case 0x50:
		// B = B
		fmt.Println("MOV B, D")
	case 0x51:
		// B = B
		fmt.Println("MOV C, D")
	case 0x52:
		// B = B
		fmt.Println("MOV D, D")
	case 0x53:
		// B = B
		fmt.Println("MOV E, D")
	case 0x54:
		// B = B
		fmt.Println("MOV H, D")
	case 0x55:
		// B = B
		fmt.Println("MOV L, D")
	case 0x56:
		// B = B
		fmt.Println("MOV (HL), D")
	case 0x57:
		// B = B
		fmt.Println("MOV A, D")
	case 0x58:
		// B = B
		fmt.Println("MOV B, E")
	case 0x59:
		// B = B
		fmt.Println("MOV C, E")
	case 0x5a:
		// B = B
		fmt.Println("MOV D, E")
	case 0x5b:
		// B = B
		fmt.Println("MOV E, E")
	case 0x5c:
		// B = B
		fmt.Println("MOV H, E")
	case 0x5d:
		// B = B
		fmt.Println("MOV L, E")
	case 0x5e:
		// B = B
		fmt.Println("MOV (HL), E")
	case 0x5f:
		// B = B
		fmt.Println("MOV A, E")
	case 0x60:
		// B = B
		fmt.Println("MOV B, H")
	case 0x61:
		// B = B
		fmt.Println("MOV C, H")
	case 0x62:
		// B = B
		fmt.Println("MOV D, H")
	case 0x63:
		// B = B
		fmt.Println("MOV E, H")
	case 0x64:
		// B = B
		fmt.Println("MOV H, H")
	case 0x65:
		// B = B
		fmt.Println("MOV L, H")
	case 0x66:
		// B = B
		fmt.Println("MOV (HL), H")
	case 0x67:
		// B = B
		fmt.Println("MOV A, H")
	case 0x68:
		// B = B
		fmt.Println("MOV B, L")
	case 0x69:
		// B = B
		fmt.Println("MOV C, L")
	case 0x6a:
		// B = B
		fmt.Println("MOV D, L")
	case 0x6b:
		// B = B
		fmt.Println("MOV E, L")
	case 0x6c:
		// B = B
		fmt.Println("MOV H, L")
	case 0x6d:
		// B = B
		fmt.Println("MOV L, L")
	case 0x6e:
		// B = B
		fmt.Println("MOV (HL), L")
	case 0x6f:
		// B = B
		fmt.Println("MOV A, L")
	case 0x70:
		// B = B
		fmt.Println("MOV B, (HL)")
	case 0x71:
		// B = B
		fmt.Println("MOV C, (HL)")
	case 0x72:
		// B = B
		fmt.Println("MOV D, (HL)")
	case 0x73:
		// B = B
		fmt.Println("MOV E, (HL)")
	case 0x74:
		// B = B
		fmt.Println("MOV H, (HL)")
	case 0x75:
		// B = B
		fmt.Println("MOV L, (HL)")
	case 0x76:
		// HALT
		fmt.Println("HALT")
	case 0x77:
		// B = B
		fmt.Println("MOV A, (HL)")
	case 0x78:
		// B = B
		fmt.Println("MOV B, A")
	case 0x79:
		// B = B
		fmt.Println("MOV C, A")
	case 0x7a:
		// B = B
		fmt.Println("MOV D, A")
	case 0x7b:
		// B = B
		fmt.Println("MOV E, A")
	case 0x7c:
		// B = B
		fmt.Println("MOV H, A")
	case 0x7d:
		// B = B
		fmt.Println("MOV L, A")
	case 0x7e:
		// B = B
		fmt.Println("MOV (HL), A")
	case 0x7f:
		// B = B
		fmt.Println("MOV A, A")
	case 0x80:
		// A = A + B
		fmt.Println("ADD B")
	case 0x81:
		// A = A + C
		fmt.Println("ADD C")
	case 0x82:
		// A = A + D
		fmt.Println("ADD D")
	case 0x83:
		// A = A + E
		fmt.Println("ADD E")
	case 0x84:
		// A = A + H
		fmt.Println("ADD H")
	case 0x85:
		// A = A + L
		fmt.Println("ADD L")
	case 0x86:
		// A = A + (HL)
		fmt.Println("ADD (HL)")
	case 0x87:
		// A = A + A
		fmt.Println("ADD A")
	case 0x88:
		// A = A + B + Cy
		fmt.Println("ADC B")
	case 0x89:
		// A = A + B + Cy
		fmt.Println("ADC C")
	case 0x8a:
		// A = A + B + Cy
		fmt.Println("ADC D")
	case 0x8b:
		// A = A + B + Cy
		fmt.Println("ADC E")
	case 0x8c:
		// A = A + B + Cy
		fmt.Println("ADC H")
	case 0x8d:
		// A = A + B + Cy
		fmt.Println("ADC L")
	case 0x8e:
		// A = A + B + Cy
		fmt.Println("ADC (HL)")
	case 0x8f:
		// A = A + B + Cy
		fmt.Println("ADC A")
	case 0x90:
		// A = A - B
		fmt.Println("SUB B")
	case 0x91:
		// A = A - B
		fmt.Println("SUB C")
	case 0x92:
		// A = A - B
		fmt.Println("SUB D")
	case 0x93:
		// A = A - B
		fmt.Println("SUB E")
	case 0x94:
		// A = A - B
		fmt.Println("SUB H")
	case 0x95:
		// A = A - B
		fmt.Println("SUB L")
	case 0x96:
		// A = A - B
		fmt.Println("SUB (HL)")
	case 0x97:
		// A = A - B
		fmt.Println("SUB A")
	case 0x98:
		// A = A - B - CY
		fmt.Println("SBB B")
	case 0x99:
		// A = A - B
		fmt.Println("SBB C")
	case 0x9a:
		// A = A - B
		fmt.Println("SBB D")
	case 0x9b:
		// A = A - B
		fmt.Println("SBB E")
	case 0x9c:
		// A = A - B
		fmt.Println("SBB H")
	case 0x9d:
		// A = A - B
		fmt.Println("SBB L")
	case 0x9e:
		// A = A - B
		fmt.Println("SBB (HL)")
	case 0x9f:
		// A = A - B
		fmt.Println("SBB A")
	case 0xa0:
		// A = A & B
		fmt.Println("AND B")
	case 0xa1:
		// A = A & B
		fmt.Println("AND C")
	case 0xa2:
		// A = A & B
		fmt.Println("AND D")
	case 0xa3:
		// A = A & B
		fmt.Println("AND E")
	case 0xa4:
		// A = A & B
		fmt.Println("AND H")
	case 0xa5:
		// A = A & B
		fmt.Println("AND L")
	case 0xa6:
		// A = A & B
		fmt.Println("AND (HL")
	case 0xa7:
		// A = A & B
		fmt.Println("AND A")
	case 0xa8:
		// A = A ^ B
		fmt.Println("XOR B")
	case 0xa9:
		// A = A ^ B
		fmt.Println("XOR C")
	case 0xaa:
		// A = A ^ B
		fmt.Println("XOR D")
	case 0xab:
		// A = A ^ B
		fmt.Println("XOR E")
	case 0xac:
		// A = A ^ B
		fmt.Println("XOR H")
	case 0xad:
		// A = A ^ B
		fmt.Println("XOR L")
	case 0xae:
		// A = A ^ B
		fmt.Println("XOR (HL)")
	case 0xaf:
		// A = A ^ B
		fmt.Println("XOR A")
	case 0xb0:
		// A = A | B
		fmt.Println("OR B")
	case 0xb1:
		// A = A | B
		fmt.Println("OR C")
	case 0xb2:
		// A = A | B
		fmt.Println("OR D")
	case 0xb3:
		// A = A | B
		fmt.Println("OR E")
	case 0xb4:
		// A = A | B
		fmt.Println("OR H")
	case 0xb5:
		// A = A | B
		fmt.Println("OR L")
	case 0xb6:
		// A = A | (HL)
		fmt.Println("OR (HL)")
	case 0xb7:
		// A = A | A
		fmt.Println("OR A")
	case 0xb8:
		// A - B
		fmt.Println("CMP B")
	case 0xb9:
		// A - B
		fmt.Println("CMP C")
	case 0xba:
		// A - B
		fmt.Println("CMP D")
	case 0xbb:
		// A - B
		fmt.Println("CMP E")
	case 0xbc:
		// A - B
		fmt.Println("CMP H")
	case 0xbd:
		// A - B
		fmt.Println("CMP L")
	case 0xbe:
		// A - B
		fmt.Println("CMP (HL)")
	case 0xbf:
		// A - B
		fmt.Println("CMP A")
	case 0xc0:
		// RET not Z
		fmt.Println("RET NZ")
	case 0xc1:
		// POP BC (C = SP; B = SP+1; SP += 2)
		fmt.Println("POP BC")
	case 0xc2:
		// JNZ addr
		fmt.Println("JNZ addr")
		opsize = 3
	case 0xc3:
		// JMP addr
		fmt.Println("JMP addr")
		opsize = 3
	case 0xc4:
		// CALL NZ
		fmt.Println("CALL NZ, addr")
		opsize = 3
	case 0xc5:
		// PUSH BC (SP-1 = B, SP-2 = C; SP = SP-2)
		fmt.Println("PUSH BC")
	case 0xc6:
		// A = A + imm
		fmt.Println("ADD D8")
		opsize = 2
	case 0xc7:
		// RST 0 - CALL $0
		fmt.Println("RST 0")
	case 0xc8:
		// RETZ
		fmt.Println("RETZ")
	case 0xc9:
		// RET PC = (SP), SP += 2
		fmt.Println("RET")
	case 0xca:
		// JZ
		fmt.Println("JZ addr")
		opsize = 3
	case 0xcc:
		// CALL Z
		fmt.Println("CALLZ addr")
		opsize = 3
	case 0xcd:
		// CALL
		fmt.Println("CALL addr")
		opsize = 3
	case 0xce:
		// ADC imm
		fmt.Println("ADC D8")
		opsize = 2
	case 0xcf:
		// RST 1 - CALL $8
		fmt.Println("RST 1")
	case 0xd0:
		// RET NC
		fmt.Println("RET NC")
	case 0xd1:
		// POP DE
		fmt.Println("POP DE")
	case 0xd2:
		// JMP NC
		fmt.Println("JMP NC, addr")
		opsize = 3
	case 0xd3:
		// OUT D8 - special
		fmt.Println("OUT D8")
		opsize = 2
	case 0xd4:
		// CALL NC
		fmt.Println("CALL NC addr")
		opsize = 3
	case 0xd5:
		// PUSH DE
		fmt.Println("PUSH DE")
	case 0xd6:
		// A = A - imm
		fmt.Println("SUB imm")
		opsize = 2
	case 0xd7:
		// RST 2 - CALL $10
		fmt.Println("RST 2")
	case 0xd8:
		// RET C
		fmt.Println("RET C")
	case 0xda:
		// JMP C
		fmt.Println("JMP C addr")
		opsize = 3
	case 0xdb:
		// In D8
		fmt.Println("IN D8")
		opsize = 2
	case 0xdc:
		// CALL C, addr
		fmt.Println("CALL C, addr")
		opsize = 3
	case 0xde:
		// A = A - imm - Cy
		fmt.Println("SBC imm")
		opsize = 2
	case 0xdf:
		// RST 3 - call $18
		fmt.Println("RST 3")
	case 0xe0:
		// RET PO
		fmt.Println("RET PO")
	case 0xe1:
		// POP HL
		fmt.Println("POP HL")
	case 0xe2:
		// JMP PO
		fmt.Println("JMP PO, addr")
		opsize = 3
	case 0xe3:
		// 	L <-> (SP); H <-> (SP+1)
		fmt.Println("XTHL")
	case 0xe4:
		// CALL PO
		fmt.Println("CALL PO, addr")
		opsize = 3
	case 0xe5:
		// PUSH HL
		fmt.Println("PUSH HL")
	case 0xe6:
		// A = A & imm
		fmt.Println("AND imm")
		opsize = 2
	case 0xe7:
		// RST 4 - call $20
		fmt.Println("RST 4")
	case 0xe8:
		// RET PE
		fmt.Println("RET PE")
	case 0xe9:
		// PCHL
		fmt.Println("PCHL")
	case 0xea:
		// JMP PE, addr
		fmt.Println("JMP PE, addr")
		opsize = 3
	case 0xeb:
		// 	H <-> D; L <-> E
		fmt.Println("XCHG")
	case 0xec:
		// CALL PE
		fmt.Println("CALL PE, addr")
		opsize = 3
	case 0xee:
		// A = A & imm
		fmt.Println("XOR imm")
		opsize = 2
	case 0xef:
		// RST 5 - CALL $28
		fmt.Println("RST 5")
	case 0xf0:
		// RET P
		fmt.Println("RET P")
	case 0xf1:
		// POP Flags
		fmt.Println("POP FLAGS")
	case 0xf2:
		// JMP P, addr
		fmt.Println("JMP P, addr")
		opsize = 3
	case 0xf3:
		// DI
		fmt.Println("DI")
	case 0xf4:
		// CALL P
		fmt.Println("CALL P, addr")
		opsize = 3
	case 0xf5:
		// PUSH FLAGS
		fmt.Println("PUSH FLAGS")
	case 0xf6:
		// A = A | imm
		fmt.Println("OR imm")
		opsize = 2
	case 0xf7:
		// RST 6 - call $30
		fmt.Println("RST 6")
	case 0xf8:
		// RET M
		fmt.Println("RET M")
	case 0xf9:
		// SPHL - SP = HL
		fmt.Println("SPHL")
	case 0xfa:
		// JMP M, addr
		fmt.Println("JMP M, addr")
		opsize = 3
	case 0xfb:
		// EI
		fmt.Println("EI")
	case 0xfc:
		// CALL M, addr
		fmt.Println("CALL M, addr")
		opsize = 3
	case 0xfe:
		// CPI - A - data
		fmt.Println("CPI data")
		opsize = 2
	case 0xff:
		// RST 7 - CALL $38
		fmt.Println("RST 7")
	default:
		// unknown op
		fmt.Printf("Uknown opcode 0x%02x\n", op)
	}

	return opsize
}
