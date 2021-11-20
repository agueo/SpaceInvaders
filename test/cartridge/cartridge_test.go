package test_cartridge

import (
	"testing"

	"github.com/agueo/SpaceInvaders/internal/cartridge"
)

func TestCartridgeRead8(t *testing.T) {
	cart, err := cartridge.NewCart("/home/hermes/go/src/github.com/agueo/SpaceInvaders/roms/invaders")
	if err != nil {
		panic(err)
	}

	tests := []struct {
		addr uint16
		want byte
	}{
		{0x00, 0x00},
		{0x03, 0xc3},
	}

	for _, test := range tests {
		result := cart.Read8(test.addr)
		if result != test.want {
			t.Errorf("Failed: wanted: %v, got: %v", test.want, result)
		}
	}
}

func TestCartridgeRead16(t *testing.T) {
	cart, err := cartridge.NewCart("/home/hermes/go/src/github.com/agueo/SpaceInvaders/roms/invaders")
	if err != nil {
		panic(err)
	}

	tests := []struct {
		addr uint16
		want []byte
	}{
		{0x00, []byte{0x00, 0x00}},
		{0x02, []byte{0x00, 0xc3}},
	}

	for _, test := range tests {
		result := cart.Read16(test.addr)
		if result[0] != test.want[0] && result[1] != test.want[1] {
			t.Errorf("Failed: wanted: %v, got: %v", test.want, result)
		}
	}
}
