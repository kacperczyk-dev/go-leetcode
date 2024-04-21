package main

import "testing"

func TestReverseWords(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world   ", "world hello"},
		{"a good   example", "example good a"},
	}
	for _, c := range cases {
		got := reverseWords(c.in)
		if got != c.want {
			t.Errorf("reverseWords(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
