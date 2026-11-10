package hack

import (
	"github.com/a-skua/nand2tetris/hack/builtin"
	"github.com/a-skua/nand2tetris/hack/chip"
)

// CPU is the Hack CPU.
type CPU struct {
	A  chip.Register
	D  chip.Register
	PC chip.PC
}

// Clock executes instruction, then updates the A, D and PC registers.
//
// inM is the value of M (RAM[A]). outM is written to addressM if writeM is 1.
// addressM and pc are the values of A and PC before the update.
// If reset is 1, PC is set to 0.
func (c *CPU) Clock(instruction, inM [16]builtin.Pin, reset builtin.Pin) (outM [16]builtin.Pin, writeM builtin.Pin, addressM, pc [16]builtin.Pin) {
	isC := instruction[15]
	isA := chip.NOT(isC)
	a := c.A.Bits()

	// comp: D op (A or M)
	y := chip.MUX16(a, inM, chip.AND(isC, instruction[12]))
	out, zr, ng := chip.ALU(
		c.D.Bits(), y,
		instruction[11], instruction[10], instruction[9],
		instruction[8], instruction[7], instruction[6],
	)

	// dest
	loadA := chip.OR(isA, chip.AND(isC, instruction[5]))
	loadD := chip.AND(isC, instruction[4])
	writeM = chip.AND(isC, instruction[3])

	// jump
	pos := chip.NOT(chip.OR(zr, ng))
	jump := chip.AND(isC, chip.OR(
		chip.OR(chip.AND(instruction[2], ng), chip.AND(instruction[1], zr)),
		chip.AND(instruction[0], pos),
	))

	outM = out
	addressM = c.A.Clock(chip.MUX16(instruction, out, isC), loadA)
	c.D.Clock(out, loadD)
	pc = c.PC.Clock(a, jump, builtin.ON, reset)
	return
}
