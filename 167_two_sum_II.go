package main

func twoSumII(numbers []int, target int) []int {
	l := len(numbers)
	ii := false
	for i, j := 0, 1; i < l && j < l; {
		n1 := numbers[i]
		n2 := numbers[j]
		if n1+n2 == target {
			return []int{i + 1, j + 1}
		}
		if n1+n2 < target {
			if ii {
				i++
			} else {
				if j+1 == l {
					i++
				} else {
					j++
				}
			}
		} else {
			j--
			ii = true
		}
	}
	return []int{}
}
