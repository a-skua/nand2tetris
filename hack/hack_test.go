package hack

import (
	"strconv"
	"testing"

	"github.com/a-skua/nand2tetris/hack/chip/builtin"
)

// word returns v as a 16-bit value (index 0 is the LSB).
func word(v int) (out [16]builtin.Pin) {
	for i := range out {
		out[i] = builtin.Pin(v >> i & 1)
	}
	return
}

func addr(v int) [15]builtin.Pin {
	w := word(v)
	return [15]builtin.Pin(w[0:15])
}

// inst parses a Hack machine code line such as "1110110000010000".
func inst(t *testing.T, s string) [16]builtin.Pin {
	t.Helper()
	v, err := strconv.ParseUint(s, 2, 16)
	if err != nil {
		t.Fatal(err)
	}
	return word(int(v))
}

func TestCPU(t *testing.T) {
	var cpu CPU
	tests := []struct {
		name         string
		instruction  string
		inM          int
		reset        builtin.Pin
		wantOutM     int
		wantWriteM   builtin.Pin
		wantAddressM int
		wantPC       int
		wantA, wantD int
	}{
		{name: "@3", instruction: "0000000000000011", wantAddressM: 0, wantPC: 0, wantA: 3},
		{name: "D=A", instruction: "1110110000010000", wantOutM: 3, wantAddressM: 3, wantPC: 1, wantA: 3, wantD: 3},
		{name: "D=D+M", instruction: "1111000010010000", inM: 4, wantOutM: 7, wantAddressM: 3, wantPC: 2, wantA: 3, wantD: 7},
		{name: "M=D-1", instruction: "1110001110001000", wantOutM: 6, wantWriteM: 1, wantAddressM: 3, wantPC: 3, wantA: 3, wantD: 7},
		{name: "AD=-1", instruction: "1110111010110000", wantOutM: 0xFFFF, wantAddressM: 3, wantPC: 4, wantA: 0xFFFF, wantD: 0xFFFF},
		{name: "@10", instruction: "0000000000001010", wantAddressM: 0xFFFF, wantPC: 5, wantA: 10, wantD: 0xFFFF},
		{name: "D;JGT (no jump)", instruction: "1110001100000001", wantOutM: 0xFFFF, wantAddressM: 10, wantPC: 6, wantA: 10, wantD: 0xFFFF},
		{name: "D;JLT (jump)", instruction: "1110001100000100", wantOutM: 0xFFFF, wantAddressM: 10, wantPC: 7, wantA: 10, wantD: 0xFFFF},
		{name: "0;JEQ (jump)", instruction: "1110101010000010", wantAddressM: 10, wantPC: 10, wantA: 10, wantD: 0xFFFF},
		{name: "reset", instruction: "0000000000000000", reset: 1, wantAddressM: 10, wantPC: 10, wantA: 0, wantD: 0xFFFF},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			instruction := inst(t, tt.instruction)
			outM, writeM, addressM, pc := cpu.Clock(instruction, word(tt.inM), tt.reset)
			if instruction[15] == builtin.OFF {
				outM = word(tt.wantOutM) // outM is undefined for A-instructions.
			}
			if outM != word(tt.wantOutM) || writeM != tt.wantWriteM || addressM != word(tt.wantAddressM) || pc != word(tt.wantPC) {
				t.Errorf("CPU.Clock() = (%v, %d, %v, %v); want (%v, %d, %v, %v)",
					outM, writeM, addressM, pc, word(tt.wantOutM), tt.wantWriteM, word(tt.wantAddressM), word(tt.wantPC))
			}
			if a, d := cpu.A.Bits(), cpu.D.Bits(); a != word(tt.wantA) || d != word(tt.wantD) {
				t.Errorf("A, D = %v, %v; want %v, %v", a, d, word(tt.wantA), word(tt.wantD))
			}
		})
	}
	if pc := cpu.PC.Bits(); pc != word(0) {
		t.Errorf("PC after reset = %v; want zero", pc)
	}
}

func TestMemory(t *testing.T) {
	var m Memory
	m.Keyboard.Clock(word('K'), builtin.ON)

	writes := []struct{ address, value int }{
		{0x0000, 1}, {0x3FFF, 2}, {0x4000, 3}, {0x5FFF, 4},
		{0x6000, 5}, // Keyboard is read-only
	}
	for _, w := range writes {
		m.Clock(word(w.value), builtin.ON, addr(w.address))
	}

	reads := []struct{ address, want int }{
		{0x0000, 1}, {0x3FFF, 2}, {0x4000, 3}, {0x5FFF, 4}, {0x6000, 'K'},
		{0x0001, 0}, {0x2000, 0}, {0x4001, 0},
	}
	for _, r := range reads {
		if got := m.Clock(word(0), builtin.OFF, addr(r.address)); got != word(r.want) {
			t.Errorf("Memory[%#x] = %v; want %v", r.address, got, word(r.want))
		}
	}
}

// prog parses Hack machine code lines.
func prog(t *testing.T, lines ...string) [][16]builtin.Pin {
	t.Helper()
	words := make([][16]builtin.Pin, len(lines))
	for i, s := range lines {
		words[i] = inst(t, s)
	}
	return words
}

// run executes cycles instructions and returns their addresses.
func run(c *Computer, cycles int) (pcs [][16]builtin.Pin) {
	for pc := range c.Clock() {
		pcs = append(pcs, pc)
		if len(pcs) >= cycles {
			break
		}
	}
	return
}

func TestNew(t *testing.T) {
	c := New(prog(t,
		"1110111111010000", // D=1
	))
	// The reset must not execute ROM[0].
	if pc, a, d := c.PC.Bits(), c.A.Bits(), c.D.Bits(); pc != word(0) || a != word(0) || d != word(0) {
		t.Errorf("after New: PC, A, D = %v, %v, %v; want all zero", pc, a, d)
	}

	run(c, 1)
	if pc, d := c.PC.Bits(), c.D.Bits(); pc != word(1) || d != word(1) {
		t.Errorf("after Clock: PC, D = %v, %v; want %v, %v", pc, d, word(1), word(1))
	}
}

func TestComputerReset(t *testing.T) {
	c := New(prog(t,
		"1110111111010000", // 0: D=1
		"0000000000000101", // 1: @5
		"1110111111010000", // 2: D=1
	))
	run(c, 2)

	c.Reset()
	if pc, a, d := c.PC.Bits(), c.A.Bits(), c.D.Bits(); pc != word(0) || a != word(0) || d != word(1) {
		t.Errorf("after Reset: PC, A, D = %v, %v, %v; want %v, %v, %v", pc, a, d, word(0), word(0), word(1))
	}
}

func TestComputerAdd(t *testing.T) {
	c := New(prog(t,
		"0000000000000010", // @2
		"1110110000010000", // D=A
		"0000000000000011", // @3
		"1110000010010000", // D=D+A
		"0000000000000000", // @0
		"1110001100001000", // M=D
	))
	pcs := run(c, 6)

	for i, pc := range pcs {
		if pc != word(i) {
			t.Errorf("cycle %d: executed %v; want %v", i, pc, word(i))
		}
	}
	if got := c.Memory.Clock(word(0), builtin.OFF, addr(0)); got != word(5) {
		t.Errorf("RAM[0] = %v; want %v", got, word(5))
	}
}

func TestComputerMax(t *testing.T) {
	program := prog(t,
		"0000000000000000", // 0: @R0
		"1111110000010000", // 1: D=M
		"0000000000000001", // 2: @R1
		"1111010011010000", // 3: D=D-M
		"0000000000001010", // 4: @10
		"1110001100000001", // 5: D;JGT
		"0000000000000001", // 6: @R1
		"1111110000010000", // 7: D=M
		"0000000000001100", // 8: @12
		"1110101010000111", // 9: 0;JMP
		"0000000000000000", // 10: @R0
		"1111110000010000", // 11: D=M
		"0000000000000010", // 12: @R2
		"1110001100001000", // 13: M=D
		"0000000000001110", // 14: @14
		"1110101010000111", // 15: 0;JMP
	)
	tests := []struct{ r0, r1, want int }{
		{r0: 3, r1: 5, want: 5},
		{r0: 9, r1: 2, want: 9},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.r0)+","+strconv.Itoa(tt.r1), func(t *testing.T) {
			c := New(program)
			c.Memory.Clock(word(tt.r0), builtin.ON, addr(0))
			c.Memory.Clock(word(tt.r1), builtin.ON, addr(1))
			// Run until the halt loop (15: 0;JMP) is reached.
			halted, cycles := false, 0
			for pc := range c.Clock() {
				if pc == word(15) {
					halted = true
					break
				}
				if cycles++; cycles >= 100 {
					break
				}
			}
			if !halted {
				t.Fatal("did not reach the halt loop in 100 cycles")
			}

			if got := c.Memory.Clock(word(0), builtin.OFF, addr(2)); got != word(tt.want) {
				t.Errorf("RAM[2] = %v; want %v", got, word(tt.want))
			}
		})
	}
}

func TestROM32K(t *testing.T) {
	var rom ROM32K
	rom.load([][16]builtin.Pin{word(10), word(11), word(12)})
	rom.Clock(word(20), builtin.ON, addr(0x4000))
	rom.Clock(word(21), builtin.ON, addr(0x7FFF))

	reads := []struct{ address, want int }{
		{0, 10}, {1, 11}, {2, 12}, {3, 0},
		{0x4000, 20}, {0x7FFF, 21}, {0x3FFF, 0},
	}
	for _, r := range reads {
		if got := rom.Clock(word(0), builtin.OFF, addr(r.address)); got != word(r.want) {
			t.Errorf("ROM32K[%#x] = %v; want %v", r.address, got, word(r.want))
		}
	}
}

func TestROM32KLoadPanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("ROM32K.load did not panic")
		}
	}()
	var rom ROM32K
	rom.load(make([][16]builtin.Pin, 1<<15+1))
}
