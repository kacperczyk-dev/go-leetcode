package main 

import (
    "reflect"
    "testing"
)

func TestRemoveDuplicatesII(t *testing.T) {
    type Want struct{
        nums []int
        l int
    }

	cases := []struct {
        in []int 
		want Want
    }{
        {[]int{1, 1, 2}, Want{[]int{1, 1, 2}, 3}},
        {[]int{7}, Want{[]int{7}, 1}},
        {[]int{-7}, Want{[]int{-7}, 1}},
        {[]int{1, 2, 3, 3, 3, 4, 5, 6, 7, 7, 7, 8, 9, 9}, Want{[]int{1, 2, 3, 3, 4, 5, 6, 7, 7, 8, 9, 9}, 12}},
        {[]int{-1, 2, 3, 3, 3, 4, 5, 6, 7, 7, 7, 8, 10}, Want{[]int{-1, 2, 3, 3, 4, 5, 6, 7, 7, 8, 10}, 11}},
    }

    for _, c := range cases {
        origNums := make([]int, len(c.in))
        copy(origNums, c.in)
		l := removeDuplicatesII(c.in)
        if !reflect.DeepEqual(c.in[:l], c.want.nums) || l != c.want.l {
            t.Errorf("removeDuplicatesII(%v) == %v, want %v", origNums, l, c.want.l)
        }
    }
}