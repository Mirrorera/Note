package d2

import "fmt"

func prevDay(n, m, d int) int {
	if n == 1 {
		return d
	}
	return prevDay(n-1, m, (d+m)*2)
}

func monkey() {
	var n, m, d int
	fmt.Scanln(&n, &m, &d)
	fmt.Println("Num:", prevDay(n, m, d))
}
