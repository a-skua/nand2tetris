package builtin

import (
	"testing"
)

func TestNAND(t *testing.T) {
	tests := []struct {
		name string
		a, b Bit
		want Bit
	}{
		{name: "NAND(0,0)", a: 0, b: 0, want: 1},
		{name: "NAND(0,1)", a: 0, b: 1, want: 1},
		{name: "NAND(1,0)", a: 1, b: 0, want: 1},
		{name: "NAND(1,1)", a: 1, b: 1, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NAND(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("NAND(%d,%d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
