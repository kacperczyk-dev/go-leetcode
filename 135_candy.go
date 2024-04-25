package main

// Two-pass approach
func candy(ratings []int) int {
	l := len(ratings)
	candies := make([]int, l)
	for i, _ := range candies {
		candies[i] = 1
	}
	for i := 0; i < l; i++ {
		prev := -1
		curr := ratings[i]
		if i > 0 {
			prev = ratings[i-1]
		}
		if prev != -1 && curr > prev {
			candies[i] = candies[i-1] + 1
		}
	}
	for i := l - 1; i >= 0; i-- {
		next := -1
		curr := ratings[i]
		if i < l-1 {
			next = ratings[i+1]
		}
		if next != -1 && curr > next && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}
	minCandies := 0
	for _, v := range candies {
		minCandies += v
	}
	return minCandies
}

// Original approach
// func candy(ratings []int) int {
// 	l := len(ratings)
// 	candies := make([]int, l)
// 	for i, _ := range candies {
// 		candies[i] = 1
// 	}
// 	for i := 0; i < l; i++ {
// 		prev := -1
// 		curr := ratings[i]
// 		next := -1
// 		if i > 0 {
// 			prev = ratings[i-1]
// 		}
// 		if i < l-1 {
// 			next = ratings[i+1]
// 		}
// 		totalAdded := 0
// 		if prev != -1 && curr > prev {
// 			k := candies[i-1] - candies[i] + 1
// 			if k > 0 {
// 				candies[i] += k
// 			}
// 			totalAdded += k
// 		}
// 		if next != -1 && curr > next {
// 			k := candies[i+1] - candies[i] + 1
// 			if k > 0 {
// 				candies[i] += k
// 			}
// 			totalAdded += k
// 		}
// 		if prev > curr && candies[i-1] <= candies[i] {
// 			i -= 2
// 		}
// 	}
// 	minCandies := 0
// 	for _, v := range candies {
// 		minCandies += v
// 	}
// 	return minCandies
// }
