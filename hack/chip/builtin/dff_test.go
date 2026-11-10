package builtin

import (
	"testing"
)

func TestDFF(t *testing.T) {
	var dff DFF
	tests := []struct {
		name     string
		load, in Bit
		want     Bit
	}{
		{name: "Clock(1, 1)", load: ON, in: ON, want: OFF},
		{name: "Clock(0, 0)", load: OFF, in: OFF, want: ON},
		{name: "Clock(0, 0)", load: OFF, in: OFF, want: ON},
		{name: "Clock(1, 0)", load: ON, in: OFF, want: ON},
		{name: "Clock(0, 0)", load: OFF, in: OFF, want: OFF},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dff.Clock(tt.load, tt.in)
			if got != tt.want {
				t.Errorf("DFF.Clock() = %v, want %v", got, tt.want)
			}
		})
	}
}
