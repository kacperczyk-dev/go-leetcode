package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = 0
	}
	return false
}
