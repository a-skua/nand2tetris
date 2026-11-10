package builtin

import "testing"

func TestNewBit(t *testing.T) {
	tests := []struct {
		name  string
		value int8
		want  Bit
	}{
		{name: "NewBit(0)", value: 0, want: 0},
		{name: "NewBit(1)", value: 1, want: 1},
		{name: "NewBit(2)", value: 2, want: 0},
		{name: "NewBit(-1)", value: -1, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBit(tt.value); got != tt.want {
				t.Errorf("NewBit() = %v, want %v", got, tt.want)
			}
		})
	}
}
