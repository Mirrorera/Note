package main

import (
	"fmt"
	"sort"
)

func main() {
	var array []int
	var pos [][]int
	var N int
	fmt.Scanln(&N)
	array = make([]int, N)
	for i, _ := range array {
		fmt.Scanf("%d", &array[i])
		pos = append(pos, make([]int, 0))
	}
	sort.Ints(array)
	maxtot := -1
	tot := 0
	pre := -1
	for _, x := range array {
		if pre != x {
			if tot > maxtot {
				maxtot = tot
			}
			pos[tot] = append(pos[tot], pre)
			pre = x
			tot = 1
		} else {
			tot++
		}
	}
	if tot > maxtot {
		maxtot = tot
	}
	pos[tot] = append(pos[tot], pre)
	fmt.Println(pos[maxtot])
}
