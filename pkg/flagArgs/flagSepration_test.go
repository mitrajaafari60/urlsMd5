package flagArgs

import (
	"testing"
)

func TestFlagsCount(t *testing.T) {
	testCase := []struct {
		input []string
		want  int
	}{
		{
			input: []string{"main.go", "-test", "10", "http://dns1.com", "http://dns2.com"},
			want:  3,
		},
		{
			input: []string{"main.go", "-test=10", "http://dns1.com", "http://dns2.com"},
			want:  2,
		},
		{
			input: []string{"main.go"},
			want:  1,
		},
	}

	for _, tc := range testCase {
		got := FlagsCount(tc.input, "test")
		if got != tc.want {
			t.Errorf("got %v, wanted %v ,%v", got, tc.want, tc.input)
		}
	}

}
