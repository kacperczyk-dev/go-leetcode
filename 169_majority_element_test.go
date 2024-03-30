package main

import (
    "testing"
)

func TestMajorityElement(t *testing.T) {
	cases := []struct {
        in []int
		want int
    }{
        {[]int{1, 1, 1, 2, 2, 2, 1}, 1},
        {[]int{9, 2, 1, 2, 2, 2, 8, 7, 2, 6, 2}, 2},
        {[]int{7}, 7},
        {[]int{-1, -1, 2, -1}, -1},
        {[]int{9, 9, 9}, 9},
    }
    for _, c := range cases {
        got := majorityElement(c.in)
        if got != c.want {
            t.Errorf("majorityElement(%v) == %v, want %v", c.in, got, c.want)
        }
    }
}