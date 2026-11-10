package builtin

import (
	"testing"
)

func TestDFF(t *testing.T) {
	var dff DFF
	tests := []struct {
		name     string
		in, want Bit
	}{
		{name: "Clock(1)", in: ON, want: OFF},
		{name: "Clock(0)", in: OFF, want: ON},
		{name: "Clock(0)", in: OFF, want: OFF},
		{name: "Clock(1)", in: ON, want: OFF},
		{name: "Clock(0)", in: OFF, want: ON},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dff.Clock(tt.in)
			if got != tt.want {
				t.Errorf("DFF.Clock() = %v, want %v", got, tt.want)
			}
		})
	}
}
