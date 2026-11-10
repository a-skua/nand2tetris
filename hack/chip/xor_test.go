package chip

import (
	"github.com/a-skua/nand2tetris/hack/builtin"
	"testing"
)

func TestXOR(t *testing.T) {
	tests := []struct {
		name string
		a, b builtin.Pin
		want builtin.Pin
	}{
		{name: "XOR(0,0)", a: 0, b: 0, want: 0},
		{name: "XOR(0,1)", a: 0, b: 1, want: 1},
		{name: "XOR(1,0)", a: 1, b: 0, want: 1},
		{name: "XOR(1,1)", a: 1, b: 1, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := XOR(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("XOR(%d,%d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
