package main

func maxProfitII(prices []int) int {
	l := len(prices)
	profit := 0
	buyPrice := prices[0]
	for i := 1; i < l; i++ {
		if prices[i] > buyPrice {
			profit += prices[i] - buyPrice

		}
		buyPrice = prices[i]
	}
	return profit
}
