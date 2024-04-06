package main

// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/description/
func removeDuplicatesII(nums []int) int {
	l := len(nums)
	ii := 1
	cnt := 1
	for i := 1; i < l; i++ {
		if nums[i] != nums[i-1] {
			nums[ii] = nums[i]
			ii++
			cnt = 1
		} else if cnt < 2 {
			nums[ii] = nums[i]
			ii++
			cnt++
		}
	}
	return ii
}

// This is already O(n)
