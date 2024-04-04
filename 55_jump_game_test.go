package main

import (
	"testing"
)

func TestCanJump(t *testing.T) {
	cases := []struct {
		in   []int
		want bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{1, 2, 3}, true},
		{[]int{3, 0, 8, 2, 0, 0, 1}, true},
		{[]int{1, 1, 1, 0}, true},
	}
	for _, c := range cases {
		got := canJump(c.in)
		if got != c.want {
			t.Errorf("majorityElement(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
