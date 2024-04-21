package main

import "strings"

func convert(s string, numRows int) string {
	str := []rune(s)
	r := make([][]rune, numRows)
	for i := range r {
		r[i] = make([]rune, len(str))
	}
	cnt := 0
	down := true
	for i, j := 0, 0; ; {
		if cnt >= len(str) {
			break
		}
		r[i][j] = str[cnt]
		if i >= numRows-1 || !down {
			down = false
			if i > 0 {
				i--
				j++
			} else {
				i++
				if i >= numRows {
					i = 0
					j++
				}
				down = true
			}
		} else {
			i++
		}
		cnt++
	}
	var sb strings.Builder
	for _, v := range r {
		for _, vv := range v {
			if vv != '\u0000' {
				sb.WriteRune(vv)
			}
		}
	}
	return sb.String()
}
