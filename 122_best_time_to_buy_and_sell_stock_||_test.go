package main

import (
	"testing"
)

func TestMaxProfitII(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 7},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1, 2, 3, 4, 5}, 4},
	}
	for _, c := range cases {
		got := maxProfitII(c.in)
		if got != c.want {
			t.Errorf("maxProfitII(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
