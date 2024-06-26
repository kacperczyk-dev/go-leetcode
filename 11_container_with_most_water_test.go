package main

import "testing"

func TestMaxArea(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}

	for _, c := range cases {
		got := maxArea(c.in)
		if got != c.want {
			t.Errorf("maxArea(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
