// Package hack implements the Hack computer.
package hack

import (
	"iter"

	"github.com/a-skua/nand2tetris/hack/chip/builtin"
)

// Computer is the Hack computer.
type Computer struct {
	CPU
	Memory
	ROM32K
}

// NewComputer returns a Computer with program stored in its ROM from address 0,
// then resets it so that execution starts from ROM[0].
// It panics if program is longer than 32768 words.
func NewComputer(program [][16]builtin.Pin) *Computer {
	c := new(Computer)
	c.ROM32K.load(program)
	c.Reset()
	return c
}

// Reset sets PC and A to 0 without executing any program instruction.
// D and the memory are not changed.
func (c *Computer) Reset() {
	var noop [16]builtin.Pin // @0
	c.CPU.Clock(noop, [16]builtin.Pin{}, builtin.ON)
}

// Clock returns an iterator that executes one instruction per iteration
// and yields the address of the executed instruction.
// It never stops by itself; break out of the loop to stop.
//
//	for pc := range c.Clock() {
//		// ...
//	}
func (c *Computer) Clock() iter.Seq[[16]builtin.Pin] {
	return func(yield func([16]builtin.Pin) bool) {
		for {
			if !yield(c.step()) {
				return
			}
		}
	}
}

// step executes one instruction and returns its address.
func (c *Computer) step() (pc [16]builtin.Pin) {
	pc, a := c.PC.Bits(), c.A.Bits()
	instruction := c.ROM32K.Clock([16]builtin.Pin{}, builtin.OFF, [15]builtin.Pin(pc[0:15]))
	inM := c.Memory.Clock([16]builtin.Pin{}, builtin.OFF, [15]builtin.Pin(a[0:15]))
	outM, writeM, addressM, _ := c.CPU.Clock(instruction, inM, builtin.OFF)
	c.Memory.Clock(outM, writeM, [15]builtin.Pin(addressM[0:15]))
	return
}
