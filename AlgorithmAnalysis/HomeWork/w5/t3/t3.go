package main

import "fmt"

var N, L int
var S []int

var dp []int

func mmin(a, b int) int {
	if b < 0 {
		return a
	}
	if a < b {
		return a
	} else {
		return b
	}
}
func main() {
	fmt.Scanln(&N, &L)
	S = make([]int, N)
	dp = make([]int, L)
	for i, _ := range S {
		fmt.Scanf("%d", &S[i])
	}
	for i, _ := range dp {
		dp[i] = -1
	}
	dp[0] = 0
	for i := 0; i < N; i++ {
		for j := S[i]; j < L; j++ {
			dp[j] = mmin(dp[j-S[i]]+1, dp[j])
		}
	}
	for i, x := range dp {
		fmt.Printf("%d: %d\n", i+1, x)
	}
}
