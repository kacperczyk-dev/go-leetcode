package main

// https://leetcode.com/problems/remove-duplicates-from-sorted-array
func removeDuplicates(nums []int) int {
	l := len(nums)
	ii := 1
	for i := 1; i < l; i++ {
		if nums[i] != nums[i-1] {
			nums[ii] = nums[i]
			ii++
		}
	}
	return ii
}

// Above is already the best solution
