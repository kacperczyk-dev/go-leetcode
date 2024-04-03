package main

// https://leetcode.com/problems/merge-sorted-array

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums := make([]int, m+n, m+n)
	i, j := 0, 0
	for i+j < m+n {
		var n1, n2 int
		ii, jj := true, true
		if i < m {
			n1 = nums1[i]
		} else {
			ii = false
		}
		if j < n {
			n2 = nums2[j]
		} else {
			jj = false
		}
		if ii && n1 < n2 {
			nums[i+j] = n1
			i++
		} else if jj && n2 < n1 {
			nums[i+j] = n2
			j++
		} else {
			if ii {
				nums[i+j] = n1
				i++
			}
			if jj {
				nums[i+j] = n2
				j++
			}
		}
	}
	copy(nums1[:], nums)
}

/*
    // Best solution, iterate backwards and add higher of the two numbers at the end of nums1
    func merge(nums1 []int, m int, nums2 []int, n int)  {
    for n != 0 {
        if m != 0 && nums1[m-1] > nums2[n-1] {
            nums1[m+n-1] = nums1[m-1]
            m--
        } else {
            nums1[m+n-1] = nums2[n-1]
            n--
        }
    }
}
*/
