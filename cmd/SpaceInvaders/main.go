package main

import (
	"fmt"

	"github.com/agueo/SpaceInvaders/internal/cartridge"
	cpu8080 "github.com/agueo/SpaceInvaders/internal/cpu"
	"github.com/agueo/SpaceInvaders/internal/disassembler"
	"github.com/agueo/SpaceInvaders/internal/memory"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	MAX_ROM uint16 = 8192
)

func main() {
	// create necessary items for emulator
	cart, err := cartridge.NewCart("/home/hermes/go/src/github.com/agueo/SpaceInvaders/roms/invaders")
	check(err)
	bus := memory.NewMemory(cart)
	cpu := new(cpu8080.Cpu)
	cpu.Init(bus)
	fmt.Println(cpu)

	var pc uint16
	for pc = 0; pc < MAX_ROM; {
		pc += disassembler.Disassemble8080Op(cart, pc)
	}
}
