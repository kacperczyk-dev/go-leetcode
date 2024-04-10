package main

// https://leetcode.com/problems/gas-station/description/
func canCompleteCircuit(gas []int, cost []int) int {
	gasAvailable := 0
	gasNeeded := 0
	possibleStartingIndexes := make([]int, 0)
	for i := 0; i < len(gas); i++ {
		gasAvailable += gas[i]
		gasNeeded += cost[i]
		if gas[i] >= cost[i] && gas[i] != 0 {
			possibleStartingIndexes = append(possibleStartingIndexes, i)
		}
	}
	if gasAvailable < gasNeeded {
		return -1
	}
	//fmt.Println(possibleStartingIndexes)
	r := -1
	lastStuckAt := -1
	for _, v := range possibleStartingIndexes {
		if v < lastStuckAt {
			continue
		}
		currGas := 0
		counter := 0
		if r != -1 {
			break
		}
		for i := v; ; i++ {
			if i >= len(gas) {
				i = 0
			}
			if gas[i] > 0 || cost[i] > 0 {
				currGas += gas[i] - cost[i]
				//fmt.Printf("Current: i = %v, g = %v, c = %v \n", i, gas[i], cost[i])
				if currGas < 0 {
					lastStuckAt = i
					break
				}
			}
			counter += 1
			if counter == len(gas) {
				r = v
				break
			}
		}
	}
	return r
}

/* Simpler solution, also done myself
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
*/
