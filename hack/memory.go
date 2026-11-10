package hack

import (
	"github.com/a-skua/nand2tetris/hack/chip"
	"github.com/a-skua/nand2tetris/hack/chip/builtin"
)

// Memory is the Hack data memory.
//
//	address       | device
//	0x0000-0x3FFF | RAM16K
//	0x4000-0x5FFF | Screen
//	0x6000        | Keyboard
type Memory struct {
	chip.RAM16K
	Screen   chip.RAM8K
	Keyboard chip.Register
}

// Clock returns the value at address, then stores in there if load is 1.
// Keyboard is read-only.
func (m *Memory) Clock(in [16]builtin.Bit, load builtin.Bit, address [15]builtin.Bit) (out [16]builtin.Bit) {
	loadRAM, loadIO := chip.DMUX(load, address[14])
	loadScreen, _ := chip.DMUX(loadIO, address[13])

	out = chip.MUX16(
		m.RAM16K.Clock(in, loadRAM, [14]builtin.Bit(address[0:14])),
		chip.MUX16(
			m.Screen.Clock(in, loadScreen, [13]builtin.Bit(address[0:13])),
			m.Keyboard.Bits(),
			address[13],
		),
		address[14],
	)
	return
}
