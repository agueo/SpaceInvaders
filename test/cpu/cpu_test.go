package test_cpu

import (
	"testing"

	cpu8080 "github.com/agueo/SpaceInvaders/internal/cpu"
)

func TestWriteToReg16(t *testing.T) {
	cpu := new(cpu8080.Cpu)

	tests := []struct {
		operands     []byte
		registerPair int
		want         error
	}{
		{[]byte{0xef, 0xbe}, cpu8080.BC, nil},
		{[]byte{0xfe, 0xca}, cpu8080.DE, nil},
		{[]byte{0xfe, 0xca}, cpu8080.HL, nil},
		{[]byte{0xef, 0xbe}, cpu8080.SP, nil},
		{[]byte{0xad, 0xde}, cpu8080.PC, nil},
	}

	for _, test := range tests {
		result := cpu.WriteReg16(test.operands, test.registerPair)
		if result != test.want {
			t.Errorf("Failed: wanted: %v, got: %v", test.want, result)
		}
	}
}

func TestReadFromReg16(t *testing.T) {
	cpu := new(cpu8080.Cpu)

	tests := []struct {
		operands []byte
		pair     int
		want     uint16
		err      error
	}{
		{[]byte{0xde, 0xaf}, cpu8080.BC, 0xafde, nil},
		{[]byte{0xfe, 0xca}, cpu8080.SP, 0xcafe, nil},
	}

	for _, test := range tests {
		cpu.WriteReg16(test.operands, test.pair)
		result, err := cpu.ReadReg16(test.pair)
		if err != nil {
			t.Errorf("failed, expecting no error")
		}
		if result != test.want {
			t.Errorf("expecting: %x, got: %x", test.want, result)
		}
	}
}
