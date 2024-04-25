package main

import "testing"

func TestStrStr(t *testing.T) {
	type In struct {
		haystack, needle string
	}
	cases := []struct {
		in   In
		want int
	}{
		{In{"sadbutsad", "sad"}, 0},
		{In{"leetcode", "leeto"}, -1},
	}

	for _, c := range cases {
		got := strStr(c.in.haystack, c.in.needle)
		if got != c.want {
			t.Errorf("strStr(%v, %v) == %v, want %v", c.in.haystack, c.in.needle, got, c.want)
		}
	}
}
