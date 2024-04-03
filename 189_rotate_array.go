package main

// https://leetcode.com/problems/rotate-array/description/

// Original solution
// func rotate(nums []int, k int)  {
//     l := len(nums)
//     kk := k % l // kk is equal to the rotation without fullcircles
//     if kk == 0 { // if kk = 0 we do only full circle(s)
//         return
//     }
//     n := make([]int, l, l)
//     copy(n[:kk], nums[l-kk:l])
//     copy(n[kk:], nums[:l-kk])
//     copy(nums, n)
// }

// Own solution: in-place, space = O(1)
func rotate(nums []int, k int) {
	l := len(nums)
	kk := k % l  // kk is equal to the rotation without fullcircles
	if kk == 0 { // if kk = 0 we do only full circle(s)
		return
	}
	n := make([]int, k, k)
	copy(n, nums[l-kk:l])
	copy(nums[kk:], nums[:l-kk])
	copy(nums[:kk], n)
}
