package main

import (
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	cases := []struct {
		s, t string
		want bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
		{"aaaaaa", "bbaaaa", false},
		{"twn", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxtxxxxxxxxxxxxxxxxxxxxwxxxxxxxxxxxxxxxxxxxxxxxxxn", true},
	}
	for _, c := range cases {
		got := isSubsequence(c.s, c.t)
		if got != c.want {
			t.Errorf("%s(%q) == %t, want %t", c.s, c.t, got, c.want)
		}
	}
}
