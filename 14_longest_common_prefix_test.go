package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	cases := []struct {
		in   []string
		want string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
	}
	for _, c := range cases {
		got := longestCommonPrefix(c.in)
		if got != c.want {
			t.Errorf("romanToInt(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
