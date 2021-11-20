package cartridge

import "io/ioutil"

type Cart struct {
	fileName string
	data     []byte
}

// Will initialize te cart rom from a file
func NewCart(fileName string) (Cart, error) {
	cart_ := Cart{fileName: fileName}
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return Cart{}, err
	}
	cart_.data = data
	return cart_, nil
}

// Read8 reads a byte from the ROM cartidge
// at the given address and returns the contents
func (c Cart) Read8(addr uint16) byte {
	return c.data[addr]
}

// Read16 reads a word from the ROM cartidge
// at the given address and returns the contents
func (c Cart) Read16(addr uint16) []byte {
	return []byte{c.data[addr], c.data[addr+1]}
}
