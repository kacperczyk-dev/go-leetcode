package main

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