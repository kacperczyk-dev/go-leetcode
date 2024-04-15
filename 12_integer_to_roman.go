package main

import (
	"math"
	"strings"
)

var intToRomanMap = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}
var nums = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

func intToRoman(num int) string {
	if r, ok := intToRomanMap[num]; ok {
		return string(r)
	}
	v := num
	r := ""
	for _, n := range nums {
		x := math.Floor(float64(v / n))
		if x >= 1 {
			r += strings.Repeat(intToRomanMap[n], int(x))
			v = v % n
		}
	}
	return r
}
