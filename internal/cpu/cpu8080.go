package cpu8080

import (
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/agueo/SpaceInvaders/internal/memory"
)

const (
	BC = iota
	DE
	HL
	PC
	SP
)

const (
	A = iota
	B
	C
	D
	E
	H
	L
)

// Registers
type Registers struct {
	A  byte
	B  byte
	C  byte
	D  byte
	E  byte
	H  byte
	L  byte
	PC [2]byte // byte 0 will be MSB byte 1 will be LSB - so big endian
	SP [2]byte
}

// Flags
type Flags struct {
	Carry    byte
	Zero     byte
	Parity   byte
	Sign     byte
	AuxCarry byte
}

// Cpu Structs
type Cpu struct {
	reg   Registers
	flags Flags
	bus   memory.Memory

	// may need some extra fields
}

func (c *Cpu) Init(bus memory.Memory) {
	c.bus = bus
}

// helper funcs for writing to Register pairs

func (c *Cpu) ReadReg8(REG int) (byte, error) {
	var value byte = 0
	var err error = nil
	switch REG {
	case A:
		value = c.reg.A
	case B:
		value = c.reg.B
	case C:
		value = c.reg.C
	case D:
		value = c.reg.D
	case E:
		value = c.reg.E
	case H:
		value = c.reg.H
	case L:
		value = c.reg.L
	}
	return value, err
}

func (c *Cpu) WriteReg8() error {
	return nil
}

// WriteReg16:
// Write a 2 byte operand to a register pair
// Parameters:
// operands: a slice of 2 bytes containing the operands bytes to be written to the register pair
// PAIR: register pair being one of BC, DE, HL, SP, PC
// Returns:
// nil on success, error on failure
func (c *Cpu) WriteReg16(operands []byte, PAIR int) error {
	if len(operands) != 2 {
		return fmt.Errorf("expecting operands slice of size 2, instead got size %d", len(operands))
	}
	switch PAIR {
	case BC:
		c.reg.B = operands[1]
		c.reg.C = operands[0]
	case DE:
		c.reg.D = operands[1]
		c.reg.E = operands[0]
	case HL:
		c.reg.H = operands[1]
		c.reg.L = operands[0]
	case SP:
		c.reg.SP[0] = operands[1]
		c.reg.SP[1] = operands[0]
	case PC:
		c.reg.PC[0] = operands[1]
		c.reg.PC[1] = operands[0]
	default:
		return fmt.Errorf("invalid pair given %d", PAIR)
	}
	return nil
}

// TODO - will have to decide if we want to return a byte array straight up
func (c *Cpu) ReadReg16(pair int) (uint16, error) {
	var ret uint16
	switch pair {
	case BC:
		data := []byte{c.reg.B, c.reg.C}
		ret = binary.BigEndian.Uint16(data)
	case DE:
		data := []byte{c.reg.D, c.reg.E}
		ret = binary.BigEndian.Uint16(data)
	case HL:
		data := []byte{c.reg.H, c.reg.L}
		ret = binary.BigEndian.Uint16(data)
	case SP:
		ret = binary.BigEndian.Uint16(c.reg.SP[:])
	case PC:
		ret = binary.BigEndian.Uint16(c.reg.PC[:])
	default:
		return 0, fmt.Errorf("invalid pair given %d", pair)
	}
	return ret, nil
}

func (c *Cpu) Read8(addr []byte) (byte, error) {
	return c.bus.Read8(addr)
}

func (c *Cpu) Write8(addr []byte, value byte) {
	if err := c.bus.Write8(addr, value); err != nil {
		panic(err)
	}
}

/**
 * String interface for printing cpu state
 */
func (c *Cpu) String() string {
	var s []string
	s = append(s, fmt.Sprintln("===== Cpu Stats ====="))
	s = append(s, fmt.Sprintln("--- Registers ---"))
	s = append(s, fmt.Sprintf("A: 0x%x\n", c.reg.A))
	s = append(s, fmt.Sprintf("B: 0x%x\n", c.reg.B))
	s = append(s, fmt.Sprintf("C: 0x%x\n", c.reg.C))
	s = append(s, fmt.Sprintf("D: 0x%x\n", c.reg.D))
	s = append(s, fmt.Sprintf("E: 0x%x\n", c.reg.E))
	s = append(s, fmt.Sprintf("H: 0x%x\n", c.reg.H))
	s = append(s, fmt.Sprintf("L: 0x%x\n", c.reg.L))
	s = append(s, fmt.Sprintln("SP:", c.reg.SP))
	s = append(s, fmt.Sprintln("PC:", c.reg.PC))
	s = append(s, fmt.Sprintln("--- Flags ---"))
	s = append(s, fmt.Sprintln("Z C N P AC"))
	s = append(s, fmt.Sprintf("%x %x %x %x %x\n", c.flags.Zero, c.flags.Carry, c.flags.Parity, c.flags.Sign, c.flags.AuxCarry))
	return strings.Join(s, "")
}
