package main

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
func maxProfit(prices []int) int {
	l := len(prices)
	max := 0
	currentMin := -1
	for i := 0; i < l; i++ {
		if prices[i] < currentMin || currentMin == -1 {
			currentMin = prices[i]
			for j := i + 1; j < l; j++ {
				profit := prices[j] - prices[i]
				if profit > max {
					max = profit
				}
			}
		}
	}
	return max
}

// Best solution - adaptation of Kodane's algorithm.
// time = O(n), space = O(1)
// func maxProfit(prices []int) int {
//     if len(prices) == 0 {
//         return 0
//     }
//     buy := prices[0]
//     profit := 0
//     for i := 1; i < len(prices); i++ {
//         if prices[i] < buy {
//             buy = prices[i]
//         } else if prices[i]-buy > profit {
//             profit = prices[i] - buy
//         }
//     }
//     return profit
// }
