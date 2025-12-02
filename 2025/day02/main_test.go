package main

import "testing"

func TestIsValid(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		// {"1", true},
		// {"998", false},
		// {"11", false},
		// {"1188511885", false},
		// {"222222", false},
		// {"38593859", false},
		// {"999", false},
		// {"565656", false},
		// {"2121212121", false},
		{"123", true},
	}

	for _, tc := range tests {
		got := IsValid(tc.input)
		if got != tc.want {
			t.Errorf("got %v, want %v for %q", got, tc.want, tc.input)
		}
	}
}
