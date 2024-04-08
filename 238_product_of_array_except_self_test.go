package main

import (
	"reflect"
	"testing"
)

func TestProductExceptSelfy(t *testing.T) {
	cases := []struct {
		in   []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}
	for _, c := range cases {
		got := productExceptSelf(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("productExceptSelf(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
