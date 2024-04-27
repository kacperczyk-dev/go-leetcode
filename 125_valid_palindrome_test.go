package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
	}
	for _, c := range cases {
		got := isPalindrome(c.in)
		if got != c.want {
			t.Errorf("isPalindrome(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}
