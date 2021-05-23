package main

import (
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// using readfile since we want the whole rom in memory
	dat, err := ioutil.ReadFile("roms/invaders")
	check(err)

	for pc := 0; pc < len(dat); {
		pc += Disassemble8080Op(dat, pc)
	}

}
