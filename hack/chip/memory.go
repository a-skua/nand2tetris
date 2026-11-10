package chip

import "github.com/a-skua/nand2tetris/hack/chip/builtin"

// Bit is a 1-bit register.
type Bit struct {
	builtin.DFF
}

// Clock returns the stored bit, then stores in if load is 1.
func (b *Bit) Clock(in, load builtin.Bit) (out builtin.Bit) {
	out = b.DFF.Clock(MUX(b.Bit, in, load))
	return
}

// Register is a 16-bit register.
type Register struct {
	bits [16]Bit
}

// Bits returns the stored value without clocking.
func (r *Register) Bits() (out [16]builtin.Bit) {
	out[0] = r.bits[0].Bit
	out[1] = r.bits[1].Bit
	out[2] = r.bits[2].Bit
	out[3] = r.bits[3].Bit
	out[4] = r.bits[4].Bit
	out[5] = r.bits[5].Bit
	out[6] = r.bits[6].Bit
	out[7] = r.bits[7].Bit
	out[8] = r.bits[8].Bit
	out[9] = r.bits[9].Bit
	out[10] = r.bits[10].Bit
	out[11] = r.bits[11].Bit
	out[12] = r.bits[12].Bit
	out[13] = r.bits[13].Bit
	out[14] = r.bits[14].Bit
	out[15] = r.bits[15].Bit
	return
}

// Clock returns the stored value, then stores in if load is 1.
func (r *Register) Clock(in [16]builtin.Bit, load builtin.Bit) (out [16]builtin.Bit) {
	out[0] = r.bits[0].Clock(in[0], load)
	out[1] = r.bits[1].Clock(in[1], load)
	out[2] = r.bits[2].Clock(in[2], load)
	out[3] = r.bits[3].Clock(in[3], load)
	out[4] = r.bits[4].Clock(in[4], load)
	out[5] = r.bits[5].Clock(in[5], load)
	out[6] = r.bits[6].Clock(in[6], load)
	out[7] = r.bits[7].Clock(in[7], load)
	out[8] = r.bits[8].Clock(in[8], load)
	out[9] = r.bits[9].Clock(in[9], load)
	out[10] = r.bits[10].Clock(in[10], load)
	out[11] = r.bits[11].Clock(in[11], load)
	out[12] = r.bits[12].Clock(in[12], load)
	out[13] = r.bits[13].Clock(in[13], load)
	out[14] = r.bits[14].Clock(in[14], load)
	out[15] = r.bits[15].Clock(in[15], load)
	return
}

// RAM8 is a memory of 8 16-bit registers.
type RAM8 struct {
	registers [8]Register
}

// Clock returns the register at address, then stores in there if load is 1.
func (r *RAM8) Clock(in [16]builtin.Bit, load builtin.Bit, address [3]builtin.Bit) (out [16]builtin.Bit) {
	// Decode the address to select which register to load.
	l0, l1, l2, l3, l4, l5, l6, l7 := DMUX8WAY(load, address)

	// Use the address to select which register's output to return.
	out = MUX8WAY16(
		r.registers[0].Clock(in, l0),
		r.registers[1].Clock(in, l1),
		r.registers[2].Clock(in, l2),
		r.registers[3].Clock(in, l3),
		r.registers[4].Clock(in, l4),
		r.registers[5].Clock(in, l5),
		r.registers[6].Clock(in, l6),
		r.registers[7].Clock(in, l7),
		address,
	)
	return
}

// RAM64 is a memory of 64 16-bit registers.
type RAM64 struct {
	rams [8]RAM8
}

// Clock returns the register at address, then stores in there if load is 1.
// address[0:3] selects the register in a [RAM8] and address[3:6] selects the RAM8.
func (r *RAM64) Clock(in [16]builtin.Bit, load builtin.Bit, address [6]builtin.Bit) (out [16]builtin.Bit) {
	low, high := [3]builtin.Bit(address[0:3]), [3]builtin.Bit(address[3:6])
	l0, l1, l2, l3, l4, l5, l6, l7 := DMUX8WAY(load, high)
	out = MUX8WAY16(
		r.rams[0].Clock(in, l0, low),
		r.rams[1].Clock(in, l1, low),
		r.rams[2].Clock(in, l2, low),
		r.rams[3].Clock(in, l3, low),
		r.rams[4].Clock(in, l4, low),
		r.rams[5].Clock(in, l5, low),
		r.rams[6].Clock(in, l6, low),
		r.rams[7].Clock(in, l7, low),
		high,
	)
	return
}

// RAM512 is a memory of 512 16-bit registers.
type RAM512 struct {
	rams [8]RAM64
}

// Clock returns the register at address, then stores in there if load is 1.
// address[0:6] selects the register in a [RAM64] and address[6:9] selects the RAM64.
func (r *RAM512) Clock(in [16]builtin.Bit, load builtin.Bit, address [9]builtin.Bit) (out [16]builtin.Bit) {
	low, high := [6]builtin.Bit(address[0:6]), [3]builtin.Bit(address[6:9])
	l0, l1, l2, l3, l4, l5, l6, l7 := DMUX8WAY(load, high)
	out = MUX8WAY16(
		r.rams[0].Clock(in, l0, low),
		r.rams[1].Clock(in, l1, low),
		r.rams[2].Clock(in, l2, low),
		r.rams[3].Clock(in, l3, low),
		r.rams[4].Clock(in, l4, low),
		r.rams[5].Clock(in, l5, low),
		r.rams[6].Clock(in, l6, low),
		r.rams[7].Clock(in, l7, low),
		high,
	)
	return
}

// RAM4K is a memory of 4096 16-bit registers.
type RAM4K struct {
	rams [8]RAM512
}

// Clock returns the register at address, then stores in there if load is 1.
// address[0:9] selects the register in a [RAM512] and address[9:12] selects the RAM512.
func (r *RAM4K) Clock(in [16]builtin.Bit, load builtin.Bit, address [12]builtin.Bit) (out [16]builtin.Bit) {
	low, high := [9]builtin.Bit(address[0:9]), [3]builtin.Bit(address[9:12])
	l0, l1, l2, l3, l4, l5, l6, l7 := DMUX8WAY(load, high)
	out = MUX8WAY16(
		r.rams[0].Clock(in, l0, low),
		r.rams[1].Clock(in, l1, low),
		r.rams[2].Clock(in, l2, low),
		r.rams[3].Clock(in, l3, low),
		r.rams[4].Clock(in, l4, low),
		r.rams[5].Clock(in, l5, low),
		r.rams[6].Clock(in, l6, low),
		r.rams[7].Clock(in, l7, low),
		high,
	)
	return
}

// RAM16K is a memory of 16384 16-bit registers.
type RAM16K struct {
	rams [4]RAM4K
}

// Clock returns the register at address, then stores in there if load is 1.
// address[0:12] selects the register in a [RAM4K] and address[12:14] selects the RAM4K.
func (r *RAM16K) Clock(in [16]builtin.Bit, load builtin.Bit, address [14]builtin.Bit) (out [16]builtin.Bit) {
	low, high := [12]builtin.Bit(address[0:12]), [2]builtin.Bit(address[12:14])
	l0, l1, l2, l3 := DMUX4WAY(load, high)
	out = MUX4WAY16(
		r.rams[0].Clock(in, l0, low),
		r.rams[1].Clock(in, l1, low),
		r.rams[2].Clock(in, l2, low),
		r.rams[3].Clock(in, l3, low),
		high,
	)
	return
}

// PC is a 16-bit program counter.
type PC struct {
	Register
}

// Clock returns the current count, then updates it.
// The priority is reset > load > inc.
//
//	reset | load | inc | next
//	1     | -    | -   | 0
//	0     | 1    | -   | in
//	0     | 0    | 1   | out+1
//	0     | 0    | 0   | out
func (p *PC) Clock(in [16]builtin.Bit, load, inc, reset builtin.Bit) (out [16]builtin.Bit) {
	now := p.Bits()
	next := MUX16(now, ADD16(now, [16]builtin.Bit{1}), inc)
	next = MUX16(next, in, load)
	next = MUX16(next, [16]builtin.Bit{}, reset)
	out = p.Register.Clock(next, OR(reset, OR(load, inc)))
	return
}
