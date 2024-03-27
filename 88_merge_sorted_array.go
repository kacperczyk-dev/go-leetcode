package main

func merge(nums1 []int, m int, nums2 []int, n int)  {
    //fmt.Println(nums1)
    //fmt.Println(nums2)
    nums := make([]int, m + n, m + n)
    i, j := 0, 0
    for ; i + j < m + n; {
        //fmt.Printf("i = %v, j = %v \n", i, j)
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
        //fmt.Printf("n1 = %v, n2 = %v \n", n1, n2)
        if ii && n1 < n2 {
            //fmt.Printf("Adding n1 = %v \n", n1)
            //nums = append(nums, n1)
            nums[i + j] = n1
            i++
        } else if jj && n2 < n1 {
            //fmt.Printf("Adding n2 = %v \n", n2)
            //nums = append(nums, n2)
            nums[i + j] = n2
            j++
        } else {
            //fmt.Printf("Adding both \n")
            if ii {
                //nums = append(nums, n1)
                nums[i + j] = n1
                i++
            }
            if jj {
                //nums = append(nums, n2)
                nums[i + j] = n2
                j++
            }
        }
    }
    copy(nums1[:], nums)
}