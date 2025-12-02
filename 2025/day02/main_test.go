package main

import "testing"

func TestIsValid(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"1", true},
		{"998", true},
		{"11", false},
		{"1188511885", false},
		{"222222", false},
		{"38593859", false},
	}

	for _, tc := range tests {
		got := IsValid(tc.input)
		if got != tc.want {
			t.Errorf("got %v, want %v for %q", got, tc.want, tc.input)
		}
	}
}
