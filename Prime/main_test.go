package prime

import (
	"reflect"
	"testing"
)

func TestPrime(t *testing.T) {
	tests := []struct {
		input int
		want  []int
	}{
		{10, []int{2, 3, 5, 7}},
		{20, []int{2, 3, 5, 7, 11, 13, 17, 19}},
	}

	for _, test := range tests {
		if got := Prime(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("Prime(%d) = %v, want %v", test.input, got, test.want)
		}
	}
}
