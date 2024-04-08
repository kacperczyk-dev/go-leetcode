package main

import (
	"testing"
)

func TestHIndex(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{3, 0, 6, 1, 5}, 3},
		{[]int{1, 3, 1}, 1},
		{[]int{5}, 1},
	}
	for _, c := range cases {
		got := hIndex(c.in)
		if got != c.want {
			t.Errorf("hIndex(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
