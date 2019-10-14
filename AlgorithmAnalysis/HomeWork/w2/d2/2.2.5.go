package d2

import (
	"fmt"
)

var dp []int

func calu(N int) int {
	if dp[N] > 0 {
		return dp[N]
	}
	if N == 1 {
		dp[N] = 1
		return dp[N]
	}
	if N%2 == 1 {
		dp[N] = calu(N/2) + calu(N/2+1)
		return dp[N]
	} else {
		dp[N] = calu(N/2) + 1
		return dp[N]
	}
}

func d225() {
	var N int
	fmt.Scanln(&N)
	dp = make([]int, N+1)
	fmt.Println(calu(N))
}
