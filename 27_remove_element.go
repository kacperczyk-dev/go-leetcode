package main

// https://leetcode.com/problems/remove-element

func removeElement(nums []int, val int) int {
	idx := len(nums) - 1
	valCnt := 0
	for i, v := range nums {
		if v == val {
			for {
				if idx <= i {
					break
				}
				if nums[idx] == val {
					idx--
				} else {
					valCnt++
					nums[i] = nums[idx]
					nums[idx] = val
					idx--
					break
				}
			}
		} else {
			valCnt++
		}
	}
	return valCnt
}

/*
// Best solution. Go through each element and if it is not the one we do not want assign it to the current index
func removeElement(nums []int, val int) int {
    i := 0
    for _, num := range nums {
        if num != val {
            nums[i] = num
            i++
        }
    }
    return i
}
*/
