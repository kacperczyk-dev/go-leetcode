package main

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i]+1 >= len(nums) {
			if i == 0 {
				return true
			}
			r := findPrevious(i, nums)
			if r {
				return r
			}
		}
	}
	return false
}

func findPrevious(j int, nums []int) bool {
	for i := j - 1; i >= 0; i-- {
		if i+nums[i] >= j {
			if i == 0 {
				return true
			}
			return findPrevious(i, nums)
		}
	}
	return false
}
