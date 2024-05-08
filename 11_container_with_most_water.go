package main

func maxArea(height []int) int {
	maxArea := -1
	for i, j := 0, len(height)-1; i < len(height) && j >= 0 && j > i; {
		h := -1
		area := -1
		if height[i] > height[j] {
			h = height[j]
			area = (j - i) * h
			if area > maxArea {
				maxArea = (j - i) * h
			}
			j--
		} else {
			h = height[i]
			area = (j - i) * h
			if area > maxArea {
				maxArea = (j - i) * h
			}
			i++
		}
	}
	return maxArea
}
