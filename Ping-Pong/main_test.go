package pingpong

import (
	"reflect"
	"testing"
)

func TestPingPong(t *testing.T) {
	tests := []struct {
		input int
		want  []string
	}{
		{1, []string{"pong"}},
		{2, []string{"pong", "pong"}},
		{3, []string{"pong", "pong", "pong"}},
		{4, []string{"pong", "pong", "pong", "pong"}},
	}

	for _, test := range tests {
		if got := PingPong(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("PingPong(%d) = %v, want %v", test.input, got, test.want)
		}
	}
}
