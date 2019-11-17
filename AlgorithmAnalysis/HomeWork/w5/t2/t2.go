package main

import "fmt"

const MAXN = 10
const MAXM = 10

var incr = []int{1, 0, 1}
var incc = []int{0, 1, 1}
var dir = []byte{'D', 'R', 'O'}

var dp [MAXN][MAXM]int
var ans [MAXN][MAXM]int
var data [MAXN][MAXM]int
var N, M int

func inmap(r int, c int) bool {
	if r >= N || c >= M || r < 0 || c < 0 {
		return false
	} else {
		return true
	}
}
func mmin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func dfs(r, c int) int {
	if dp[r][c] >= 0 {
		return dp[r][c]
	}
	//dp[r][c] = data[r][c]
	for i := 0; i < 3; i++ {
		nr := r + incr[i]
		nc := c + incc[i]
		if !inmap(nr, nc) {
			continue
		}
		d := data[r][c] + dfs(nr, nc)
		if d < dp[r][c] || dp[r][c] < 0 {
			dp[r][c] = d
			ans[r][c] = i
		}
	}
	if dp[r][c] < 0 {
		dp[r][c] = data[r][c]
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
	/*
		for _, x := range dp {
			fmt.Println(x)
		}
		for _, x := range ans {
			fmt.Println(x)
		}
		for _, x := range ans {
			fmt.Println(string(x[0:len(x)]))
		}
	*/
	r, c := 0, 0
	var sans string
	for r < N && c < M {
		//fmt.Println(r, c)
		//fmt.Printf("%c\n", dir[ans[r][c]])
		sans += fmt.Sprintf("%c", dir[ans[r][c]])
		//直接r+= incr c+=incc 会导致c所加的值不是原处的方向
		nr := r + incr[ans[r][c]]
		nc := c + incc[ans[r][c]]
		r = nr
		c = nc
		if r == N-1 && c == M-1 {
			break
		}
	}
	fmt.Println(sans)
}
