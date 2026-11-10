// Package hack implements the Hack computer.
package hack

import "github.com/a-skua/nand2tetris/hack/chip/builtin"

// Computer is the Hack computer.
type Computer struct {
	CPU
	Memory
	ROM32K
}

// NewComputer returns a Computer with program stored in its ROM from address 0,
// then resets it so that execution starts from ROM[0].
// It panics if program is longer than 32768 words.
func NewComputer(program [][16]builtin.Bit) *Computer {
	c := new(Computer)
	c.ROM32K.load(program)
	c.Reset()
	return c
}

// Reset sets PC and A to 0 without executing any program instruction.
// D and the memory are not changed.
func (c *Computer) Reset() {
	var noop [16]builtin.Bit // @0
	c.CPU.Clock(noop, [16]builtin.Bit{}, builtin.ON)
}

// Clock executes one instruction. If reset is 1, PC is set to 0.
func (c *Computer) Clock(reset builtin.Bit) {
	pc, a := c.PC.Bits(), c.A.Bits()
	instruction := c.ROM32K.Clock([16]builtin.Bit{}, builtin.OFF, [15]builtin.Bit(pc[0:15]))
	inM := c.Memory.Clock([16]builtin.Bit{}, builtin.OFF, [15]builtin.Bit(a[0:15]))
	outM, writeM, addressM, _ := c.CPU.Clock(instruction, inM, reset)
	c.Memory.Clock(outM, writeM, [15]builtin.Bit(addressM[0:15]))
}
