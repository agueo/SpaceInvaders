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
		fmt.Printf("%04x NOP\n", PC)
	case 0x01:
		// B = byte 3, C = byte 2
		fmt.Printf("%04x LDI D16, BC\n", PC)
		opsize = 3
	case 0x02:
		// (BC) = A
		fmt.Printf("%04x STR A (BC)\n", PC)
	case 0x03:
		// BC = BC + 1
		fmt.Printf("%04x INC BC\n", PC)
	case 0x04:
		// B = B + 1
		fmt.Printf("%04x INC B\n", PC)
	case 0x05:
		// B = B - 1
		fmt.Printf("%04x DEC B\n", PC)
	case 0x06:
		// B = byte 2
		fmt.Printf("%04x MVI D8, B\n", PC)
		opsize = 2
	case 0x07:
		// 	A = A << 1; bit 0 = prev bit 7; CY = prev bit 7
		fmt.Printf("%04x RLC\n", PC)
	case 0x09:
		// HL = HL + BC
		fmt.Printf("%04x DAD BC\n", PC)
	case 0x0a:
		// A = (BC)
		fmt.Printf("%04x LD (BC) A\n", PC)
	case 0x0b:
		// BC = BC - 1
		fmt.Printf("%04x DEC BC\n", PC)
	case 0x0c:
		// C = C + 1
		fmt.Printf("%04x INC C\n", PC)
	case 0x0d:
		// C = C - 1
		fmt.Printf("%04x DEC C\n", PC)
	case 0x0e:
		// C = byte 2
		fmt.Printf("%04x MVI D8, C\n", PC)
		opsize = 2
	case 0x0f:
		// A = A >> 1 bit 7 = prev bit 0
		fmt.Printf("%04x RRC\n", PC)
	case 0x11:
		// D = byte 3, E = byte 2
		fmt.Printf("%04x LDI D16, DE\n", PC)
		opsize = 3
	case 0x12:
		// (DE) = A
		fmt.Printf("%04x STR A, (DE)\n", PC)
	case 0x13:
		// DE = DE + 1
		fmt.Printf("%04x INC DE\n", PC)
	case 0x14:
		// D = D + 1
		fmt.Printf("%04x INC D\n", PC)
	case 0x15:
		// D = D - 1
		fmt.Printf("%04x DEC D\n", PC)
	case 0x16:
		// D = byte 2
		fmt.Printf("%04x MVI D8, D\n", PC)
		opsize = 2
	case 0x17:
		// A = A << 1
		fmt.Printf("%04x RAL\n", PC)
	case 0x19:
		// HL + HL + DE
		fmt.Printf("%04x DAD DE\n", PC)
	case 0x1a:
		// A = (DE)
		fmt.Printf("%04x LD A, (DE)\n", PC)
	case 0x1b:
		// DE = DE - 1
		fmt.Printf("%04x DEC DE\n", PC)
	case 0x1c:
		// E = E + 1
		fmt.Printf("%04x INC E\n", PC)
	case 0x1d:
		// E = E - 1
		fmt.Printf("%04x DEC E\n", PC)
	case 0x1e:
		// E = byte 2
		fmt.Printf("%04x MVI D8, E\n", PC)
		opsize = 2
	case 0x1f:
		// A = A >> 1
		fmt.Printf("%04x RAR\n", PC)
	case 0x20:
		// Special
		fmt.Printf("%04x RIM\n", PC)
	case 0x21:
		// H = byte 3, L = byte 2
		fmt.Printf("%04x LDI D16, HL\n", PC)
		opsize = 3
	case 0x22:
		// (adr) = L, (adr + 1) = H
		fmt.Printf("%04x SHLD H, D16\n", PC)
		opsize = 3
	case 0x23:
		// HL = HL + 1
		fmt.Printf("%04x INC HL\n", PC)
	case 0x24:
		// H = H + 1
		fmt.Printf("%04x INC H\n", PC)
	case 0x25:
		// H = H - 1
		fmt.Printf("%04x DEC H\n", PC)
	case 0x26:
		// L = byte 2
		fmt.Printf("%04x MVI L, D8\n", PC)
		opsize = 2
	case 0x27:
		// H = byte 2
		fmt.Printf("%04x DAA\n", PC)
	case 0x29:
		// HL = HL + HL
		fmt.Printf("%04x DAD H\n", PC)
	case 0x2a:
		// L = (byte 2); H = (byte 3)
		fmt.Printf("%04x LHLD adr\n", PC)
		opsize = 3
	case 0x2b:
		// HL = HL - 1
		fmt.Printf("%04x DEC HL\n", PC)
	case 0x2c:
		// L = L + 1
		fmt.Printf("%04x INC L\n", PC)
	case 0x2d:
		// L = L - 1
		fmt.Printf("%04x DEC L\n", PC)
	case 0x2e:
		// L = byte 2
		fmt.Printf("%04x MVI L, D8\n", PC)
		opsize = 2
	case 0x2f:
		// A = !A
		fmt.Printf("%04x CMA\n", PC)
	case 0x30:
		// SIM - special
		fmt.Printf("%04x SIM\n", PC)
	case 0x31:
		// SP.hi = byte 3, SP.lo = byte 2
		fmt.Printf("%04x LDI D16, SP\n", PC)
		opsize = 3
	case 0x32:
		// (addr) = A
		fmt.Printf("%04x STA addr\n", PC)
		opsize = 3
	case 0x33:
		// SP = SP + 1
		fmt.Printf("%04x INC SP\n", PC)
	case 0x34:
		// (HL) = (HL) + 1
		fmt.Printf("%04x INC (HL)\n", PC)
	case 0x35:
		// (HL) = (HL) - 1
		fmt.Printf("%04x DEC (HL)\n", PC)
	case 0x36:
		// (HL) = byte 2
		fmt.Printf("%04x MVI D8 (HL)\n", PC)
		opsize = 2
	case 0x37:
		// CY = 1
		fmt.Printf("%04x STC\n", PC)
	case 0x39:
		// HL = HL + SP
		fmt.Printf("%04x DAD SP\n", PC)
	case 0x3a:
		// A = (Addr)
		fmt.Printf("%04x LDA adr\n", PC)
		opsize = 3
	case 0x3b:
		// SP = SP - 1
		fmt.Printf("%04x DEC SP\n", PC)
	case 0x3c:
		// A = A + 1
		fmt.Printf("%04x INC A\n", PC)
	case 0x3d:
		// A = A - 1
		fmt.Printf("%04x DEC A\n", PC)
	case 0x3e:
		// A = byte 2
		fmt.Printf("%04x MVI D8, A\n", PC)
		opsize = 2
	case 0x3f:
		// CY != CY
		fmt.Printf("%04x CMP C\n", PC)
	case 0x40:
		// B = B
		fmt.Printf("%04x MOV B, B\n", PC)
	case 0x41:
		// B = C
		fmt.Printf("%04x MOV C, B\n", PC)
	case 0x42:
		// B = D
		fmt.Printf("%04x MOV D, B\n", PC)
	case 0x43:
		// B = H
		fmt.Printf("%04x MOV E, B\n", PC)
	case 0x44:
		// B = B
		fmt.Printf("%04x MOV H, B\n", PC)
	case 0x45:
		// B = B
		fmt.Printf("%04x MOV L, B\n", PC)
	case 0x46:
		// B = B
		fmt.Printf("%04x MOV (HL), B\n", PC)
	case 0x47:
		// B = B
		fmt.Printf("%04x MOV A, B\n", PC)
	case 0x48:
		// B = B
		fmt.Printf("%04x MOV B, C\n", PC)
	case 0x49:
		// B = B
		fmt.Printf("%04x MOV C, C\n", PC)
	case 0x4a:
		// B = B
		fmt.Printf("%04x MOV D, C\n", PC)
	case 0x4b:
		// B = B
		fmt.Printf("%04x MOV E, C\n", PC)
	case 0x4c:
		// B = B
		fmt.Printf("%04x MOV H, C\n", PC)
	case 0x4d:
		// B = B
		fmt.Printf("%04x MOV L, C\n", PC)
	case 0x4e:
		// B = B
		fmt.Printf("%04x MOV (HL), C\n", PC)
	case 0x4f:
		// B = B
		fmt.Printf("%04x MOV A, C\n", PC)
	case 0x50:
		// B = B
		fmt.Printf("%04x MOV B, D\n", PC)
	case 0x51:
		// B = B
		fmt.Printf("%04x MOV C, D\n", PC)
	case 0x52:
		// B = B
		fmt.Printf("%04x MOV D, D\n", PC)
	case 0x53:
		// B = B
		fmt.Printf("%04x MOV E, D\n", PC)
	case 0x54:
		// B = B
		fmt.Printf("%04x MOV H, D\n", PC)
	case 0x55:
		// B = B
		fmt.Printf("%04x MOV L, D\n", PC)
	case 0x56:
		// B = B
		fmt.Printf("%04x MOV (HL), D\n", PC)
	case 0x57:
		// B = B
		fmt.Printf("%04x MOV A, D\n", PC)
	case 0x58:
		// B = B
		fmt.Printf("%04x MOV B, E\n", PC)
	case 0x59:
		// B = B
		fmt.Printf("%04x MOV C, E\n", PC)
	case 0x5a:
		// B = B
		fmt.Printf("%04x MOV D, E\n", PC)
	case 0x5b:
		// B = B
		fmt.Printf("%04x MOV E, E\n", PC)
	case 0x5c:
		// B = B
		fmt.Printf("%04x MOV H, E\n", PC)
	case 0x5d:
		// B = B
		fmt.Printf("%04x MOV L, E\n", PC)
	case 0x5e:
		// B = B
		fmt.Printf("%04x MOV (HL), E\n", PC)
	case 0x5f:
		// B = B
		fmt.Printf("%04x MOV A, E\n", PC)
	case 0x60:
		// B = B
		fmt.Printf("%04x MOV B, H\n", PC)
	case 0x61:
		// B = B
		fmt.Printf("%04x MOV C, H\n", PC)
	case 0x62:
		// B = B
		fmt.Printf("%04x MOV D, H\n", PC)
	case 0x63:
		// B = B
		fmt.Printf("%04x MOV E, H\n", PC)
	case 0x64:
		// B = B
		fmt.Printf("%04x MOV H, H\n", PC)
	case 0x65:
		// B = B
		fmt.Printf("%04x MOV L, H\n", PC)
	case 0x66:
		// B = B
		fmt.Printf("%04x MOV (HL), H\n", PC)
	case 0x67:
		// B = B
		fmt.Printf("%04x MOV A, H\n", PC)
	case 0x68:
		// B = B
		fmt.Printf("%04x MOV B, L\n", PC)
	case 0x69:
		// B = B
		fmt.Printf("%04x MOV C, L\n", PC)
	case 0x6a:
		// B = B
		fmt.Printf("%04x MOV D, L\n", PC)
	case 0x6b:
		// B = B
		fmt.Printf("%04x MOV E, L\n", PC)
	case 0x6c:
		// B = B
		fmt.Printf("%04x MOV H, L\n", PC)
	case 0x6d:
		// B = B
		fmt.Printf("%04x MOV L, L\n", PC)
	case 0x6e:
		// B = B
		fmt.Printf("%04x MOV (HL), L\n", PC)
	case 0x6f:
		// B = B
		fmt.Printf("%04x MOV A, L\n", PC)
	case 0x70:
		// B = B
		fmt.Printf("%04x MOV B, (HL)\n", PC)
	case 0x71:
		// B = B
		fmt.Printf("%04x MOV C, (HL)\n", PC)
	case 0x72:
		// B = B
		fmt.Printf("%04x MOV D, (HL)\n", PC)
	case 0x73:
		// B = B
		fmt.Printf("%04x MOV E, (HL)\n", PC)
	case 0x74:
		// B = B
		fmt.Printf("%04x MOV H, (HL)\n", PC)
	case 0x75:
		// B = B
		fmt.Printf("%04x MOV L, (HL)\n", PC)
	case 0x76:
		// HALT
		fmt.Printf("%04x HALT\n", PC)
	case 0x77:
		// B = B
		fmt.Printf("%04x MOV A, (HL)\n", PC)
	case 0x78:
		// B = B
		fmt.Printf("%04x MOV B, A\n", PC)
	case 0x79:
		// B = B
		fmt.Printf("%04x MOV C, A\n", PC)
	case 0x7a:
		// B = B
		fmt.Printf("%04x MOV D, A\n", PC)
	case 0x7b:
		// B = B
		fmt.Printf("%04x MOV E, A\n", PC)
	case 0x7c:
		// B = B
		fmt.Printf("%04x MOV H, A\n", PC)
	case 0x7d:
		// B = B
		fmt.Printf("%04x MOV L, A\n", PC)
	case 0x7e:
		// B = B
		fmt.Printf("%04x MOV (HL), A\n", PC)
	case 0x7f:
		// B = B
		fmt.Printf("%04x MOV A, A\n", PC)
	case 0x80:
		// A = A + B
		fmt.Printf("%04x ADD B\n", PC)
	case 0x81:
		// A = A + C
		fmt.Printf("%04x ADD C\n", PC)
	case 0x82:
		// A = A + D
		fmt.Printf("%04x ADD D\n", PC)
	case 0x83:
		// A = A + E
		fmt.Printf("%04x ADD E\n", PC)
	case 0x84:
		// A = A + H
		fmt.Printf("%04x ADD H\n", PC)
	case 0x85:
		// A = A + L
		fmt.Printf("%04x ADD L\n", PC)
	case 0x86:
		// A = A + (HL)
		fmt.Printf("%04x ADD (HL)\n", PC)
	case 0x87:
		// A = A + A
		fmt.Printf("%04x ADD A\n", PC)
	case 0x88:
		// A = A + B + Cy
		fmt.Printf("%04x ADC B\n", PC)
	case 0x89:
		// A = A + B + Cy
		fmt.Printf("%04x ADC C\n", PC)
	case 0x8a:
		// A = A + B + Cy
		fmt.Printf("%04x ADC D\n", PC)
	case 0x8b:
		// A = A + B + Cy
		fmt.Printf("%04x ADC E\n", PC)
	case 0x8c:
		// A = A + B + Cy
		fmt.Printf("%04x ADC H\n", PC)
	case 0x8d:
		// A = A + B + Cy
		fmt.Printf("%04x ADC L\n", PC)
	case 0x8e:
		// A = A + B + Cy
		fmt.Printf("%04x ADC (HL)\n", PC)
	case 0x8f:
		// A = A + B + Cy
		fmt.Printf("%04x ADC A\n", PC)
	case 0x90:
		// A = A - B
		fmt.Printf("%04x SUB B\n", PC)
	case 0x91:
		// A = A - B
		fmt.Printf("%04x SUB C\n", PC)
	case 0x92:
		// A = A - B
		fmt.Printf("%04x SUB D\n", PC)
	case 0x93:
		// A = A - B
		fmt.Printf("%04x SUB E\n", PC)
	case 0x94:
		// A = A - B
		fmt.Printf("%04x SUB H\n", PC)
	case 0x95:
		// A = A - B
		fmt.Printf("%04x SUB L\n", PC)
	case 0x96:
		// A = A - B
		fmt.Printf("%04x SUB (HL)\n", PC)
	case 0x97:
		// A = A - B
		fmt.Printf("%04x SUB A\n", PC)
	case 0x98:
		// A = A - B - CY
		fmt.Printf("%04x SBB B\n", PC)
	case 0x99:
		// A = A - B
		fmt.Printf("%04x SBB C\n", PC)
	case 0x9a:
		// A = A - B
		fmt.Printf("%04x SBB D\n", PC)
	case 0x9b:
		// A = A - B
		fmt.Printf("%04x SBB E\n", PC)
	case 0x9c:
		// A = A - B
		fmt.Printf("%04x SBB H\n", PC)
	case 0x9d:
		// A = A - B
		fmt.Printf("%04x SBB L\n", PC)
	case 0x9e:
		// A = A - B
		fmt.Printf("%04x SBB (HL)\n", PC)
	case 0x9f:
		// A = A - B
		fmt.Printf("%04x SBB A\n", PC)
	case 0xa0:
		// A = A & B
		fmt.Printf("%04x AND B\n", PC)
	case 0xa1:
		// A = A & B
		fmt.Printf("%04x AND C\n", PC)
	case 0xa2:
		// A = A & B
		fmt.Printf("%04x AND D\n", PC)
	case 0xa3:
		// A = A & B
		fmt.Printf("%04x AND E\n", PC)
	case 0xa4:
		// A = A & B
		fmt.Printf("%04x AND H\n", PC)
	case 0xa5:
		// A = A & B
		fmt.Printf("%04x AND L\n", PC)
	case 0xa6:
		// A = A & B
		fmt.Printf("%04x AND (HL\n", PC)
	case 0xa7:
		// A = A & B
		fmt.Printf("%04x AND A\n", PC)
	case 0xa8:
		// A = A ^ B
		fmt.Printf("%04x XOR B\n", PC)
	case 0xa9:
		// A = A ^ B
		fmt.Printf("%04x XOR C\n", PC)
	case 0xaa:
		// A = A ^ B
		fmt.Printf("%04x XOR D\n", PC)
	case 0xab:
		// A = A ^ B
		fmt.Printf("%04x XOR E\n", PC)
	case 0xac:
		// A = A ^ B
		fmt.Printf("%04x XOR H\n", PC)
	case 0xad:
		// A = A ^ B
		fmt.Printf("%04x XOR L\n", PC)
	case 0xae:
		// A = A ^ B
		fmt.Printf("%04x XOR (HL)\n", PC)
	case 0xaf:
		// A = A ^ B
		fmt.Printf("%04x XOR A\n", PC)
	case 0xb0:
		// A = A | B
		fmt.Printf("%04x OR B\n", PC)
	case 0xb1:
		// A = A | B
		fmt.Printf("%04x OR C\n", PC)
	case 0xb2:
		// A = A | B
		fmt.Printf("%04x OR D\n", PC)
	case 0xb3:
		// A = A | B
		fmt.Printf("%04x OR E\n", PC)
	case 0xb4:
		// A = A | B
		fmt.Printf("%04x OR H\n", PC)
	case 0xb5:
		// A = A | B
		fmt.Printf("%04x OR L\n", PC)
	case 0xb6:
		// A = A | (HL)
		fmt.Printf("%04x OR (HL)\n", PC)
	case 0xb7:
		// A = A | A
		fmt.Printf("%04x OR A\n", PC)
	case 0xb8:
		// A - B
		fmt.Printf("%04x CMP B\n", PC)
	case 0xb9:
		// A - B
		fmt.Printf("%04x CMP C\n", PC)
	case 0xba:
		// A - B
		fmt.Printf("%04x CMP D\n", PC)
	case 0xbb:
		// A - B
		fmt.Printf("%04x CMP E\n", PC)
	case 0xbc:
		// A - B
		fmt.Printf("%04x CMP H\n", PC)
	case 0xbd:
		// A - B
		fmt.Printf("%04x CMP L\n", PC)
	case 0xbe:
		// A - B
		fmt.Printf("%04x CMP (HL)\n", PC)
	case 0xbf:
		// A - B
		fmt.Printf("%04x CMP A\n", PC)
	case 0xc0:
		// RET not Z
		fmt.Printf("%04x RET NZ\n", PC)
	case 0xc1:
		// POP BC (C = SP; B = SP+1; SP += 2)
		fmt.Printf("%04x POP BC\n", PC)
	case 0xc2:
		// JNZ addr
		fmt.Printf("%04x JNZ addr\n", PC)
		opsize = 3
	case 0xc3:
		// JMP addr
		fmt.Printf("%04x JMP addr\n", PC)
		opsize = 3
	case 0xc4:
		// CALL NZ
		fmt.Printf("%04x CALL NZ, addr\n", PC)
		opsize = 3
	case 0xc5:
		// PUSH BC (SP-1 = B, SP-2 = C; SP = SP-2)
		fmt.Printf("%04x PUSH BC\n", PC)
	case 0xc6:
		// A = A + imm
		fmt.Printf("%04x ADD D8\n", PC)
		opsize = 2
	case 0xc7:
		// RST 0 - CALL $0
		fmt.Printf("%04x RST 0\n", PC)
	case 0xc8:
		// RETZ
		fmt.Printf("%04x RETZ\n", PC)
	case 0xc9:
		// RET PC = (SP), SP += 2
		fmt.Printf("%04x RET\n", PC)
	case 0xca:
		// JZ
		fmt.Printf("%04x JZ addr\n", PC)
		opsize = 3
	case 0xcc:
		// CALL Z
		fmt.Printf("%04x CALLZ addr\n", PC)
		opsize = 3
	case 0xcd:
		// CALL
		fmt.Printf("%04x CALL addr\n", PC)
		opsize = 3
	case 0xce:
		// ADC imm
		fmt.Printf("%04x ADC D8\n", PC)
		opsize = 2
	case 0xcf:
		// RST 1 - CALL $8
		fmt.Printf("%04x RST 1\n", PC)
	case 0xd0:
		// RET NC
		fmt.Printf("%04x RET NC\n", PC)
	case 0xd1:
		// POP DE
		fmt.Printf("%04x POP DE\n", PC)
	case 0xd2:
		// JMP NC
		fmt.Printf("%04x JMP NC, addr\n", PC)
		opsize = 3
	case 0xd3:
		// OUT D8 - special
		fmt.Printf("%04x OUT D8\n", PC)
		opsize = 2
	case 0xd4:
		// CALL NC
		fmt.Printf("%04x CALL NC addr\n", PC)
		opsize = 3
	case 0xd5:
		// PUSH DE
		fmt.Printf("%04x PUSH DE\n", PC)
	case 0xd6:
		// A = A - imm
		fmt.Printf("%04x SUB imm\n", PC)
		opsize = 2
	case 0xd7:
		// RST 2 - CALL $10
		fmt.Printf("%04x RST 2\n", PC)
	case 0xd8:
		// RET C
		fmt.Printf("%04x RET C\n", PC)
	case 0xda:
		// JMP C
		fmt.Printf("%04x JMP C addr\n", PC)
		opsize = 3
	case 0xdb:
		// In D8
		fmt.Printf("%04x IN D8\n", PC)
		opsize = 2
	case 0xdc:
		// CALL C, addr
		fmt.Printf("%04x CALL C, addr\n", PC)
		opsize = 3
	case 0xde:
		// A = A - imm - Cy
		fmt.Printf("%04x SBC imm\n", PC)
		opsize = 2
	case 0xdf:
		// RST 3 - call $18
		fmt.Printf("%04x RST 3\n", PC)
	case 0xe0:
		// RET PO
		fmt.Printf("%04x RET PO\n", PC)
	case 0xe1:
		// POP HL
		fmt.Printf("%04x POP HL\n", PC)
	case 0xe2:
		// JMP PO
		fmt.Printf("%04x JMP PO, addr\n", PC)
		opsize = 3
	case 0xe3:
		// 	L <-> (SP); H <-> (SP+1)
		fmt.Printf("%04x XTHL\n", PC)
	case 0xe4:
		// CALL PO
		fmt.Printf("%04x CALL PO, addr\n", PC)
		opsize = 3
	case 0xe5:
		// PUSH HL
		fmt.Printf("%04x PUSH HL\n", PC)
	case 0xe6:
		// A = A & imm
		fmt.Printf("%04x AND imm\n", PC)
		opsize = 2
	case 0xe7:
		// RST 4 - call $20
		fmt.Printf("%04x RST 4\n", PC)
	case 0xe8:
		// RET PE
		fmt.Printf("%04x RET PE\n", PC)
	case 0xe9:
		// PCHL
		fmt.Printf("%04x PCHL\n", PC)
	case 0xea:
		// JMP PE, addr
		fmt.Printf("%04x JMP PE, addr\n", PC)
		opsize = 3
	case 0xeb:
		// 	H <-> D; L <-> E
		fmt.Printf("%04x XCHG\n", PC)
	case 0xec:
		// CALL PE
		fmt.Printf("%04x CALL PE, addr\n", PC)
		opsize = 3
	case 0xee:
		// A = A & imm
		fmt.Printf("%04x XOR imm\n", PC)
		opsize = 2
	case 0xef:
		// RST 5 - CALL $28
		fmt.Printf("%04x RST 5\n", PC)
	case 0xf0:
		// RET P
		fmt.Printf("%04x RET P\n", PC)
	case 0xf1:
		// POP Flags
		fmt.Printf("%04x POP FLAGS\n", PC)
	case 0xf2:
		// JMP P, addr
		fmt.Printf("%04x JMP P, addr\n", PC)
		opsize = 3
	case 0xf3:
		// DI
		fmt.Printf("%04x DI\n", PC)
	case 0xf4:
		// CALL P
		fmt.Printf("%04x CALL P, addr\n", PC)
		opsize = 3
	case 0xf5:
		// PUSH FLAGS
		fmt.Printf("%04x PUSH FLAGS\n", PC)
	case 0xf6:
		// A = A | imm
		fmt.Printf("%04x OR imm\n", PC)
		opsize = 2
	case 0xf7:
		// RST 6 - call $30
		fmt.Printf("%04x RST 6\n", PC)
	case 0xf8:
		// RET M
		fmt.Printf("%04x RET M\n", PC)
	case 0xf9:
		// SPHL - SP = HL
		fmt.Printf("%04x SPHL\n", PC)
	case 0xfa:
		// JMP M, addr
		fmt.Printf("%04x JMP M, addr\n", PC)
		opsize = 3
	case 0xfb:
		// EI
		fmt.Printf("%04x EI\n", PC)
	case 0xfc:
		// CALL M, addr
		fmt.Printf("%04x CALL M, addr\n", PC)
		opsize = 3
	case 0xfe:
		// CPI - A - data
		fmt.Printf("%04x CPI data\n", PC)
		opsize = 2
	case 0xff:
		// RST 7 - CALL $38
		fmt.Printf("%04x RST 7\n", PC)
	default:
		// unknown op
		fmt.Printf("Uknown opcode 0x%02x\n", op)
	}

	return opsize
}
