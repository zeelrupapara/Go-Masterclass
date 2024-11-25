package reverse

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"hello", "olleh"},
		{"world", "dlrow"},
	}

	for _, test := range tests {
		if got := Reverse(test.input); got != test.want {
			t.Errorf("Reverse(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}
