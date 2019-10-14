package d2

import "fmt"

var N int
var A []int

func divideArray() ([]int, []int) {
	var l, r int = 0, N - 1
	for l < r {
		for A[l] >= 0 {
			l++
		}
		for A[r] < 0 {
			r--
		}
		if l >= r {
			break
		}
		swap(&A[l], &A[r])
	}
	B := A[0:l]
	C := A[l:N]
	return B, C
}
func DivideArray() {
	fmt.Scanln(&N)
	A = make([]int, N)
	for i, _ := range A {
		fmt.Scanf("%d", &A[i])
	}
	B, C := divideArray()

	fmt.Println("B:", B)
	fmt.Println("C:", C)
}

func swap(l *int, r *int) {
	t := *l
	*l = *r
	*r = t
}
