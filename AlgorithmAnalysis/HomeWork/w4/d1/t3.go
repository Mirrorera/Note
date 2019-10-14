package d1

import "fmt"

var N int
var array [][]int
var aim int

func sqrBinSearch(lr int, lc int, rr int, rc int) (exist bool) {
	if lr < 0 || lr >= N || lc < 0 || lc > N || rr < 0 || rr >= N || rc < 0 || rc >= N {
		return false
	}
	if lr > rr || lc > rc {
		return false
	}
	if lr == rr && lc == rc {
		if array[lr][lc] == aim {
			return true
		} else {
			return false
		}
	}
	mr := (lr + rr) >> 1
	mc := (lc + rc) >> 1
	d := array[mr][mc]
	if aim < d {
		exist = sqrBinSearch(lr, lc, mr-1, mc-1)
		if exist {
			return
		}
		exist = sqrBinSearch(mr, lc, rr, mc-1)
		if exist {
			return
		}
		exist = sqrBinSearch(lr, mc, mr-1, rc)
		if exist {
			return
		}
	} else if aim > d {
		exist = sqrBinSearch(mr+1, mc+1, rr, rc)
		if exist {
			return
		}
		exist = sqrBinSearch(mr+1, lc, rr, mc)
		if exist {
			return
		}
		exist = sqrBinSearch(lr, mc+1, mr, rc)
		if exist {
			return
		}
	} else {
		return true
	}
	return false
}

func main() {
	fmt.Scanln(&N)
	for i := 0; i < N; i++ {
		t := make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Scanf("%d", &t[j])
		}
		array = append(array, t)
		fmt.Scanf("\n")
	}
	fmt.Scanln(&aim)
	//fmt.Println(array, aim)
	ans := sqrBinSearch(0, 0, N-1, N-1)
	if ans {
		fmt.Println("Exist")
	} else {
		fmt.Println("Not Exist")
	}
}
