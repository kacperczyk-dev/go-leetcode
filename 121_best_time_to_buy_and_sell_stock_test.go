package main

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}
	for _, c := range cases {
		got := maxProfit(c.in)
		if got != c.want {
			t.Errorf("maxProfit(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
