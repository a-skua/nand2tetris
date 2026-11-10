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

func TestNewWord(t *testing.T) {
	tests := []struct {
		name  string
		value int16
		want  Word
	}{
		{name: "NewWord(0)", value: 0, want: Word{}},
		{name: "NewWord(1)", value: 1, want: Word{1}},
		{name: "NewWord(6)", value: 6, want: Word{0, 1, 1}},
		{name: "NewWord(0x7FFF)", value: 0x7FFF, want: Word{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0}},
		{name: "NewWord(-1)", value: -1, want: Word{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
		{name: "NewWord(-0x8000)", value: -0x8000, want: Word{15: 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWord(tt.value); got != tt.want {
				t.Errorf("NewWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
