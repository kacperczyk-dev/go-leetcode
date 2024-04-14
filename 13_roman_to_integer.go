package main

func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	l := len(s)
	sum := 0
	for i := l - 1; i >= 0; i-- {
		currVal := m[s[i]]
		if i != l-1 && m[s[i+1]] > currVal {
			sum -= currVal
		} else {
			sum += currVal
		}
	}
	return sum
}
