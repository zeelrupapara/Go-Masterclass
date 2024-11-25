package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{3, "Fizz"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
		{7, "7"},
	}

	for _, test := range tests {
		if got := FizzBuzz(test.input); got != test.want {
			t.Errorf("FizzBuzz(%d) = %q, want %q", test.input, got, test.want)
		}
	}
}
