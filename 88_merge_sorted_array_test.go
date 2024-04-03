package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type In struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	cases := []struct {
		in   In
		want []int
	}{
		{In{[]int{1, 4, 9, 0, 0, 0, 0}, 3, []int{2, 8, 9, 12}, 4}, []int{1, 2, 4, 8, 9, 9, 12}},
		{In{[]int{2, 4, 9, 0, 0, 0, 0}, 3, []int{1, 3, 9, 12}, 4}, []int{1, 2, 3, 4, 9, 9, 12}},
		{In{[]int{-3, 1, 4, 9, 0, 0, 0, 0}, 4, []int{2, 8, 9, 12}, 4}, []int{-3, 1, 2, 4, 8, 9, 9, 12}},
		{In{[]int{0, 0, 0, 0}, 0, []int{2, 8, 9, 12}, 4}, []int{2, 8, 9, 12}},
		{In{[]int{2, 8, 9, 12}, 4, []int{0, 0, 0, 0}, 0}, []int{2, 8, 9, 12}},
	}
	for _, c := range cases {
		merge(c.in.nums1, c.in.m, c.in.nums2, c.in.n)
		got := c.in.nums1
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("merge(%v, %v, %v, %v) == %v, want %v", c.in.nums1, c.in.m, c.in.nums2, c.in.n, got, c.want)
		}
	}
}
