package d1

import "fmt"

var N int
var cnt int = 0

func f311() {
	fmt.Scanln(&N)
	for i := 1; i <= N; i++ {
		n := i
		for n%5 == 0 {
			cnt++
			n /= 5
		}
	}
	fmt.Println("Ans =", cnt)
}
