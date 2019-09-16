package main

import (
	"fmt"
)

var N int
var data []int

//var ans int = 0

func getLBound(l, r int) (ll int) {
	//fmt.Println(l, r)

	for l <= r {
		mid := (l + r) >> 1
		if mid == data[mid] {
			ll = getLBound(l, mid-1)
			if ll == -1 {
				ll = mid
			}
			return
		}
		if mid > data[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	ll = -1
	return
}
func getRBound(l, r int) (rr int) {
	//fmt.Println(l, r)

	for l <= r {
		mid := (l + r) >> 1
		if mid == data[mid] {
			rr = getRBound(mid+1, r)
			if rr == -1 {
				rr = mid
			}
			return
		}
		if mid > data[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	rr = -1
	return
}

func main() {
	fmt.Scanln(&N)
	data = make([]int, N+1)
	for i, _ := range data[1 : N+1] {
		fmt.Scanf("%d", &data[i+1])
	}
	//fmt.Println(data[1 : N+1])
	//mid := (N + 1) / 2
	l, r := getLBound(1, N), getRBound(1, N)

	//ans = r - l + 1
	fmt.Printf("ANS = [%d, %d]", l, r)
}
