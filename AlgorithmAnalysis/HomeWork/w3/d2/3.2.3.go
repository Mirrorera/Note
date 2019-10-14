package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d\n", &N)
	var A, B []int
	A = make([]int, N)
	B = make([]int, N)
	for i, _ := range A {
		fmt.Scanf("%d", &A[i])
	}
	//fmt.Println(A)
	for i, _ := range A {
		if i != 0 {
			B[i] = B[i-1] * A[i-1]
		} else {
			B[i] = 1
		}
	}
	L := len(A)
	for i, _ := range A {
		if i != 0 {
			A[L-i-1] *= A[L-i]
		}
	}

	for i, _ := range B {
		if i != L-1 {
			B[i] = B[i] * A[i+1]
		}
	}

	fmt.Println(B)
}
