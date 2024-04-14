package main

import "testing"

func TestRomanToInt(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"V", 5},
		{"MD", 1500},
		{"IX", 9},
		{"IV", 4},
		{"XVII", 17},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"XLIX", 49},
	}
	for _, c := range cases {
		got := romanToInt(c.in)
		if got != c.want {
			t.Errorf("romanToInt(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
