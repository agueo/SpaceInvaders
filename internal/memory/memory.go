package memory

import (
	"encoding/binary"
	"fmt"

	"github.com/agueo/SpaceInvaders/internal/cartridge"
)

// Memory Map
// 0000-1FFF 8K ROM
// 2000-23FF 1K RAM
// 2400-3FFF 7K Video RAM
// 4000- RAM mirror
type Memory struct {
	cart cartridge.Cart
	ram  []byte
	vram []byte
}

func NewMemory(cart cartridge.Cart) Memory {
	m := Memory{}
	m.cart = cart
	m.ram = make([]byte, 0x400)
	m.vram = make([]byte, 0x1c00)
	return m
}

func (m *Memory) Read8(addr []byte) (byte, error) {
	addr_ := binary.BigEndian.Uint16(addr)
	var value byte
	var err error = nil
	if addr_ <= 0x1FFF {
		value = m.cart.Read8(addr_)
	} else if addr_ >= 0x2000 && addr_ <= 0x23FF {
		value = m.ram[addr_-0x2000]
	} else if addr_ >= 0x2400 && addr_ <= 0x3FFF {
		value = m.vram[addr_]
	} else {
		err = fmt.Errorf("tried to read from invalid address: %x", addr)
	}
	return value, err
}

func (m *Memory) Write8(addr []byte, data byte) error {
	addr_ := binary.BigEndian.Uint16(addr)
	var err error = nil

	if addr_ >= 0x2000 && addr_ <= 0x23FF {
		m.ram[addr_-0x2000] = data
	} else if addr_ >= 0x2400 && addr_ <= 0x3FFF {
		m.vram[addr_] = data
	} else {
		err = fmt.Errorf("tried to read from invalid address: %x", addr)
	}
	return err
}
