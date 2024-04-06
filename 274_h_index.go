package main

// https://leetcode.com/problems/h-index/description/
func hIndex(citations []int) int {
	if len(citations) == 1 && citations[0] > 0 {
		return 1
	}
	max := 0
	m := make(map[int]int)
	for _, v := range citations {
		if v > max {
			max = v
		}
		m[v] += 1
	}
	sum := 0
	for c := max; c > 0; c-- {
		articles, ok := m[c]
		if !ok {
			continue
		}
		if articles+sum <= c {
			sum += articles
		} else {
			if sum < c {
				sum = c
			}
		}
	}
	return sum
}
