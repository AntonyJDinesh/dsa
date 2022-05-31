package dp

func minCostClimbingStairs(cost []int) int {
	if cost == nil {
		return 0
	}

	if len(cost) == 1 {
		return cost[0]
	}

	twoStep := 0
	oneStep := 0

	// 1, 2, 1
	for i := 2; i < len(cost)+1; i++ {
		temp := oneStep
		oneStep = min(oneStep+cost[i-1], twoStep+cost[i-2])
		twoStep = temp
	}

	return oneStep
}

func min(val1, val2 int) int {
	if val1 < val2 {
		return val1
	}

	return val2
}
