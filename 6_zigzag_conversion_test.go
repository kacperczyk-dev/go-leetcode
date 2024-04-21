package main

import "testing"

func TestConvert(t *testing.T) {
	type In struct {
		s       string
		numRows int
	}
	cases := []struct {
		in   In
		want string
	}{
		{In{"PAYPALISHIRING", 3}, "PAHNAPLSIIGYIR"},
		{In{"PAYPALISHIRING", 4}, "PINALSIGYAHRPI"},
		{In{"A", 1}, "A"},
	}
	for _, c := range cases {
		got := convert(c.in.s, c.in.numRows)
		if got != c.want {
			t.Errorf("convert(%q, %q) == %q, want %q", c.in.s, c.in.numRows, got, c.want)
		}
	}
}
