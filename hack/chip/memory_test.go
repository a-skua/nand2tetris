package chip

import (
	"github.com/a-skua/nand2tetris/hack/chip/builtin"
	"testing"
)

func TestBit(t *testing.T) {
	var bit Bit
	tests := []struct {
		name     string
		load, in builtin.Pin
		want     builtin.Pin
	}{
		{name: "Clock(1,1)", in: 1, load: 1, want: 0},
		{name: "Clock(0,0)", in: 0, load: 0, want: 1},
		{name: "Clock(0,0)", in: 0, load: 0, want: 1},
		{name: "Clock(0,1)", in: 0, load: 1, want: 1},
		{name: "Clock(0,0)", in: 0, load: 1, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bit.Clock(tt.in, tt.load)
			if got != tt.want {
				t.Errorf("Bit.Clock(%d,%d) = %d; want %d", tt.load, tt.in, got, tt.want)
			}
		})
	}
}

func TestRegister(t *testing.T) {
	var reg Register
	tests := []struct {
		name string
		in   [16]builtin.Pin
		load builtin.Pin
		want [16]builtin.Pin
	}{
		{
			name: "Clock",
			in:   [16]builtin.Pin{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0}, load: 1,
			want: [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "Clock",
			in:   [16]builtin.Pin{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1}, load: 0,
			want: [16]builtin.Pin{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
		},
		{
			name: "Clock",
			in:   [16]builtin.Pin{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1}, load: 0,
			want: [16]builtin.Pin{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
		},
		{
			name: "Clock",
			in:   [16]builtin.Pin{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1}, load: 1,
			want: [16]builtin.Pin{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
		},
		{
			name: "Clock",
			in:   [16]builtin.Pin{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1}, load: 0,
			want: [16]builtin.Pin{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reg.Clock(tt.in, tt.load)
			if got != tt.want {
				t.Errorf("Register.Clock(%v,%d) = %v; want %v", tt.in, tt.load, got, tt.want)
			}
		})
	}
}

// toBits returns the lowest n bits of v (index 0 is the LSB).
func toBits(v, n int) []builtin.Pin {
	out := make([]builtin.Pin, n)
	for i := range out {
		out[i] = builtin.Pin(v >> i & 1)
	}
	return out
}

// word returns v as a 16-bit value.
func word(v int) [16]builtin.Pin {
	return [16]builtin.Pin(toBits(v, 16))
}

// testRAM writes a distinct value to each address and reads them back.
func testRAM(t *testing.T, clock func(in [16]builtin.Pin, load builtin.Pin, address int) [16]builtin.Pin, addresses []int) {
	t.Helper()
	value := func(address int) [16]builtin.Pin { return word(address ^ 0xA5A5) }

	for _, a := range addresses {
		if got := clock(value(a), 1, a); got != word(0) {
			t.Errorf("write %d: Clock() = %v; want zero", a, got)
		}
	}
	for _, a := range addresses {
		// load=0 must not overwrite the stored value.
		for range 2 {
			if got := clock(word(0xFFFF), 0, a); got != value(a) {
				t.Errorf("read %d: Clock() = %v; want %v", a, got, value(a))
			}
		}
	}
}

func TestRAM8(t *testing.T) {
	var ram RAM8
	testRAM(t, func(in [16]builtin.Pin, load builtin.Pin, a int) [16]builtin.Pin {
		return ram.Clock(in, load, [3]builtin.Pin(toBits(a, 3)))
	}, []int{0, 1, 2, 3, 4, 5, 6, 7})
}

func TestRAM64(t *testing.T) {
	var ram RAM64
	testRAM(t, func(in [16]builtin.Pin, load builtin.Pin, a int) [16]builtin.Pin {
		return ram.Clock(in, load, [6]builtin.Pin(toBits(a, 6)))
	}, []int{0, 1, 7, 8, 9, 42, 63})
}

func TestRAM512(t *testing.T) {
	var ram RAM512
	testRAM(t, func(in [16]builtin.Pin, load builtin.Pin, a int) [16]builtin.Pin {
		return ram.Clock(in, load, [9]builtin.Pin(toBits(a, 9)))
	}, []int{0, 1, 63, 64, 65, 300, 511})
}

func TestRAM4K(t *testing.T) {
	var ram RAM4K
	testRAM(t, func(in [16]builtin.Pin, load builtin.Pin, a int) [16]builtin.Pin {
		return ram.Clock(in, load, [12]builtin.Pin(toBits(a, 12)))
	}, []int{0, 1, 511, 512, 513, 2000, 4095})
}

func TestRAM16K(t *testing.T) {
	var ram RAM16K
	testRAM(t, func(in [16]builtin.Pin, load builtin.Pin, a int) [16]builtin.Pin {
		return ram.Clock(in, load, [14]builtin.Pin(toBits(a, 14)))
	}, []int{0, 1, 4095, 4096, 4097, 10000, 16383})
}

func TestPC(t *testing.T) {
	var pc PC
	tests := []struct {
		name             string
		in               int
		load, inc, reset builtin.Pin
		want             int
	}{
		{name: "inc", inc: 1, want: 0},
		{name: "inc", inc: 1, want: 1},
		{name: "hold", want: 2},
		{name: "load", in: 100, load: 1, want: 2},
		{name: "load>inc", in: 200, load: 1, inc: 1, want: 100},
		{name: "reset>load", in: 300, load: 1, inc: 1, reset: 1, want: 200},
		{name: "hold", want: 0},
		{name: "load -1", in: 0xFFFF, load: 1, want: 0},
		{name: "inc overflow", inc: 1, want: 0xFFFF},
		{name: "hold", want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pc.Clock(word(tt.in), tt.load, tt.inc, tt.reset)
			if got != word(tt.want) {
				t.Errorf("PC.Clock(%d,%d,%d,%d) = %v; want %v", tt.in, tt.load, tt.inc, tt.reset, got, word(tt.want))
			}
		})
	}
}

func TestRAM8K(t *testing.T) {
	var ram RAM8K
	testRAM(t, func(in [16]builtin.Pin, load builtin.Pin, a int) [16]builtin.Pin {
		return ram.Clock(in, load, [13]builtin.Pin(toBits(a, 13)))
	}, []int{0, 1, 4095, 4096, 4097, 6000, 8191})
}
