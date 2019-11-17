package main

import "fmt"

const MAXN = 500 + 13
const MAXM = 21 + 2

var N, M int
var data []int

var dp [MAXN][MAXM]float64
var pre [MAXN][MAXM]int

var d [MAXN][MAXN]float64

func mmin(a, b float64) float64 {
	if a < b {
		return a
	} else {
		return b
	}
}
func main() {
	fmt.Scanln(&N, &M)
	data = make([]int, N)
	for i, _ := range data {
		fmt.Scanf("%d", &data[i])
		dp[i][0] = float64(data[i])
	}
	fmt.Scanln()

	for l, _ := range data {
		d[l][l] += float64(data[l])
		for r := l + 1; r < N; r++ {
			d[l][r] = d[l][r-1]*10 + float64(data[r])
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j <= M+1; j++ {
			dp[i][j] = -1
			pre[i][j] = -1
		}
	}
	for i := 0; i < N; i++ { //r
		for j := 1; j <= M+1; j++ { //m
			for k := 0; k <= i; k++ { //l
				add := 0.0
				if k > 0 {
					if dp[k-1][j-1] < 0 {
						continue
					}
					add += dp[k-1][j-1]
				}
				if dp[i][j] < 0 || d[k][i]+add < dp[i][j] {
					pre[i][j] = k - 1
					dp[i][j] = d[k][i] + add
				}
				//dp[i][j] = mmin(dp[i][j], d[k][i]+add)
			}
		}
	}

	var pos []int = make([]int, N)
	r, c := N-1, M+1
	for c > 0 && pre[r][c] >= 0 {
		pos[pre[r][c]] = 1
		r = pre[r][c]
		c--
	}
	//fmt.Println(pos)
	for i, x := range data {
		fmt.Printf("%c", byte(int('0')+x))
		if pos[i] == 1 {
			fmt.Print("+")
		}
	}
	fmt.Println(" =", dp[N-1][M+1])
}
