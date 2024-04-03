package main

// https://leetcode.com/problems/majority-element/description/
// time = O(n), space O(n)
func majorityElement(nums []int) int {
	l := len(nums)
	m := make(map[int]int)
	r := 0
	for _, v := range nums {
		m[v] += 1
		if m[v] > l/2 {
			r = v
			break
		}
	}
	return r
}

// Best solution: time = O(n), space = O(n)
// func majorityElement(nums []int) int {
//     var ans int
//     var count int
//     for _, num := range nums {
//         if count == 0 {
//             ans = num
//         }
//         if num == ans {
//             count++
//         } else {
//             count--
//         }
//     }
//     return ans
// }
