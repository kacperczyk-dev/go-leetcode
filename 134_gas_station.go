package main

// https://leetcode.com/problems/gas-station/description/
func canCompleteCircuit(gas []int, cost []int) int {
	startingAt := 0
	currentGas := 0
	counter := 0
	i := 0
	for {
		if counter == len(gas) {
			break
		}
		if i >= len(gas) {
			i = 0
		}
		if currentGas+gas[i] < cost[i] {
			if i < startingAt {
				startingAt = -1
				break
			}
			startingAt = i + 1
			currentGas = 0
			counter = 0
		} else {
			currentGas += gas[i] - cost[i]
			counter++
		}
		i++
	}
	return startingAt
}
