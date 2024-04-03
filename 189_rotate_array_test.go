package main

import (
	"reflect"
	"testing"
)

func TestRotateArray(t *testing.T) {
	type In struct {
		arr []int
		k   int
	}
	cases := []struct {
		in   In
		want []int
	}{
		{In{[]int{1, 2, 3, 4, 5, 6, 7}, 3}, []int{5, 6, 7, 1, 2, 3, 4}},
		{In{[]int{1, 2, 3, 4, 5, 6, 7}, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
		{In{[]int{1, 2, 3, 4, 5, 6, 7}, 49}, []int{1, 2, 3, 4, 5, 6, 7}},
		{In{[]int{1}, 35}, []int{1}},
	}
	for _, c := range cases {
		orig := make([]int, len(c.in.arr), len(c.in.arr))
		copy(orig, c.in.arr)
		rotate(c.in.arr, c.in.k)
		got := c.in.arr
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("merge(%v, %v) == %v, want %v", c.in.arr, c.in.k, got, c.want)
		}
	}
}
