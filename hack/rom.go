package hack

import (
	"github.com/a-skua/nand2tetris/hack/builtin"
	"github.com/a-skua/nand2tetris/hack/chip"
)

// ROM32K is the Hack instruction memory of 32768 16-bit words.
type ROM32K [2]chip.RAM16K

// load stores program from address 0.
// It panics if program is longer than 32768 words.
func (r *ROM32K) load(program [][16]builtin.Pin) {
	if len(program) > 1<<15 {
		panic("hack: program is longer than 32768 words")
	}
	var address [16]builtin.Pin
	for _, instruction := range program {
		r.Clock(instruction, builtin.ON, [15]builtin.Pin(address[0:15]))
		address = chip.ADD16(address, [16]builtin.Pin{1})
	}
}

// Clock returns the word at address, then stores in there if load is 1.
// address[0:14] selects the word in a [chip.RAM16K] and address[14] selects the RAM16K.
func (r *ROM32K) Clock(in [16]builtin.Pin, load builtin.Pin, address [15]builtin.Pin) (out [16]builtin.Pin) {
	low := [14]builtin.Pin(address[0:14])
	l0, l1 := chip.DMUX(load, address[14])
	out = chip.MUX16(
		r[0].Clock(in, l0, low),
		r[1].Clock(in, l1, low),
		address[14],
	)
	return
}
