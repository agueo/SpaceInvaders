package test_disassembler

import (
	"testing"

	"github.com/agueo/SpaceInvaders/internal/cartridge"
	"github.com/agueo/SpaceInvaders/internal/disassembler"
)

func TestDisassemble8080Op(t *testing.T) {
	cart, err := cartridge.NewCart("/home/hermes/go/src/github.com/agueo/SpaceInvaders/roms/invaders")
	if err != nil {
		panic(err)
	}
	result := disassembler.Disassemble8080Op(cart, 0)
	if result != 1 {
		t.Errorf("Failed to read the rom properly")
	}
}
