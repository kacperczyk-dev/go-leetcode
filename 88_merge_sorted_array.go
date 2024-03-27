package main

func merge(nums1 []int, m int, nums2 []int, n int)  {
    nums := make([]int, m + n, m + n)
    i, j := 0, 0
    for ; i + j < m + n; {
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
            nums[i + j] = n1
            i++
        } else if jj && n2 < n1 {
            nums[i + j] = n2
            j++
        } else {
            if ii {
                nums[i + j] = n1
                i++
            }
            if jj {
                nums[i + j] = n2
                j++
            }
        }
    }
    copy(nums1[:], nums)
}