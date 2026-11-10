package builtin

import "testing"

func TestNewPin(t *testing.T) {
	tests := []struct {
		name  string
		value int8
		want  Pin
	}{
		{name: "NewPin(0)", value: 0, want: 0},
		{name: "NewPin(1)", value: 1, want: 1},
		{name: "NewPin(2)", value: 2, want: 0},
		{name: "NewPin(-1)", value: -1, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPin(tt.value); got != tt.want {
				t.Errorf("NewPin() = %v, want %v", got, tt.want)
			}
		})
	}
}
