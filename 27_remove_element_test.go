package main

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	type In struct {
		nums []int
		val  int
	}
	type Want struct {
		nums []int
		l    int
	}
	cases := []struct {
		in   In
		want Want
	}{
		{In{[]int{1, 9, 2, 2, 3, 4, 5, 2, 9, 2}, 2}, Want{[]int{1, 9, 9, 5, 3, 4}, 6}},
		{In{[]int{1}, 1}, Want{[]int{}, 0}},
		{In{[]int{2}, 1}, Want{[]int{2}, 1}},
	}
	for _, c := range cases {
		origNums := make([]int, len(c.in.nums))
		copy(origNums, c.in.nums)
		l := removeElement(c.in.nums, c.in.val)
		if !reflect.DeepEqual(c.in.nums[:l], c.want.nums) || l != c.want.l {
			t.Errorf("removeElement(%v, %v) == %v, want %v", origNums, c.in.val, c.in.nums[:l], c.want.nums)
		}
	}
}
