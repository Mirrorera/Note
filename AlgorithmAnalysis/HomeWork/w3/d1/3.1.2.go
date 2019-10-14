package main

import "fmt"

var graph [][]int
var M, N int

var incr = []int{1, 0, -1, 0}
var incc = []int{0, 1, 0, -1}
var inc = []int{1, 1, -1, -1}
var gap []int

func move(d int, r int, c int, stat int) {
	//fmt.Println(r, c, d)
	stat %= 4
	graph[r][c] = d
	if d == M*N {
		return
	}

	var oper *int
	if incr[stat] != 0 {
		oper = &r
	} else {
		oper = &c
	}
	*oper += inc[stat]
	if *oper >= gap[stat] || *oper < 0 || graph[r][c] != 0 {
		*oper -= inc[stat]
		stat++
		move(d, r, c, stat)
	} else {
		move(d+1, r, c, stat)
	}
}
func main() {
	//var N int
	fmt.Scanln(&M, &N)
	gap = []int{M, N, M, N}
	for i := 0; i < M; i++ {
		var t []int = make([]int, N)
		graph = append(graph, t)
	}
	move(1, 0, 0, 0)
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%3d", graph[i][j])
		}
		fmt.Println()
	}
}
