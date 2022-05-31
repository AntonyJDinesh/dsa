package dp

func tribonacci(n int) int {

	prev3, prev2, prev1 := 0, 1, 1
	if n == 0 {
		return prev3
	}

	if n == 1 || n == 2 {
		return prev1
	}

	for i := 3; i <= n; i++ {
		//fmt.Printf("%d: b=> prev1: %d, prev2: %d, prev3: %d\n", i, prev1, prev2, prev3)
		prev1, prev2, prev3 = prev1+prev2+prev3, prev1, prev2
		//fmt.Printf("%d: a=> prev1: %d, prev2: %d, prev3: %d\n", i, prev1, prev2, prev3)
	}

	return prev1
}
