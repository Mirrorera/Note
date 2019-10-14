package d1

import (
	"fmt"
)

var ans int = 0

func mergeSort(array []int) {

	L := len(array)
	if len(array) == 1 {
		return
	}
	mid := L >> 1
	var oper []int = make([]int, L)
	copy(oper, array)
	ll := oper[0:mid]
	rr := oper[mid:L]
	//fmt.Println(oper, ll, rr)
	mergeSort(ll)
	mergeSort(rr)
	//fmt.Println(oper, ll, rr)
	l := 0
	r := 0
	for l < mid && r < L-mid {
		if ll[l] < rr[r] {
			array[l+r] = ll[l]
			l++
		} else {
			array[l+r] = rr[r]
			r++
			ans += mid - l
		}
	}
	for l < mid {
		array[l+r] = ll[l]
		l++
	}
	for r < L-mid {
		array[l+r] = rr[r]
		r++
	}

	//log.Println(ll, rr, oper)
}

func d1() {
	var N int
	var array []int

	fmt.Scanln(&N)
	array = make([]int, N)
	for i, _ := range array {
		fmt.Scanf("%d", &array[i])
	}

	mergeSort(array)
	fmt.Println("ANS =", ans)
}
