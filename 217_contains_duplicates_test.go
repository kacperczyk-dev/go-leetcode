package main

import "testing"

func TestContainsDuplicate(t *testing.T) {
	cases := []struct {
		in   []int
		want bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, c := range cases {
		got := containsDuplicate(c.in)
		if got != c.want {
			t.Errorf("containsDuplicate(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
