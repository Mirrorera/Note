package main

import "fmt"

const MAXN = 10
const MAXM = 10

var incr = []int{1, 0, 1}
var incc = []int{0, 1, 1}

var dp [MAXN][MAXM]int
var data [MAXN][MAXM]int
var N, M int

func inmap(r int, c int) bool {
	if r >= N || c >= M || r < 0 || c < 0 {
		return false
	} else {
		return true
	}
}
func mmax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func dfs(r, c int) int {
	if dp[r][c] >= 0 {
		return dp[r][c]
	}
	dp[r][c] = data[r][c]
	for i := 0; i < 3; i++ {
		nr := r + incr[i]
		nc := c + incc[i]
		if !inmap(nr, nc) {
			continue
		}
		dp[r][c] = mmax(dp[r][c], data[r][c]+dfs(nr, nc))
	}
	return dp[r][c]
}

func main() {
	fmt.Scanln(&N, &M)
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			fmt.Scanf("%d", &data[i][j])
			dp[i][j] = -1
		}
		fmt.Scanln()
	}
	dfs(0, 0)
	for _, x := range dp {
		fmt.Println(x)
	}
}
