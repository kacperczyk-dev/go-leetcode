package main

/*
pLeft[i] holds the product of all elements on the left of i in nums
pRight[j] holds the product of all elements on the right of j in nums
pLeft[i] * pRight[i] = product of all elements from nums except for nums[i]
*/
func productExceptSelf(nums []int) []int {
	pLeft := make([]int, len(nums))
	pRight := make([]int, len(nums))
	for i, _ := range nums {
		j := len(nums) - 1 - i
		if i == 0 {
			pLeft[i] = 1
			pRight[j] = 1
		} else {
			pLeft[i] = nums[i-1] * pLeft[i-1]
			pRight[j] = nums[j+1] * pRight[j+1]
		}
	}
	for i, _ := range nums {
		nums[i] = pLeft[i] * pRight[i]
	}
	return nums
}
