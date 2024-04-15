package main

import "testing"

func TestIntToRoman(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{5, "V"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
		{3, "III"},
		{90, "XC"},
		{576, "DLXXVI"},
	}
	for _, c := range cases {
		got := intToRoman(c.in)
		if got != c.want {
			t.Errorf("intToRoman(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
