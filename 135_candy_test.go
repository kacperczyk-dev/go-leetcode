package main

import (
	"testing"
)

func TestCandy(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 0, 2}, 5},
		{[]int{1, 2, 2}, 4},
		{[]int{1, 2, 87, 87, 87, 2, 1}, 13},
		{[]int{1, 2, 3, 1, 0}, 9},
		{[]int{1, 6, 10, 8, 7, 3, 2}, 18},
		{[]int{1, 3, 4, 5, 2}, 11},
	}
	for _, c := range cases {
		got := candy(c.in)
		if got != c.want {
			t.Errorf("candy(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
