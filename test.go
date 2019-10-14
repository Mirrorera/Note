package main

import "fmt"

func main() {
	var a []int = []int{0, 1, 2, 3}
	b := a[0:2]
	b[0] = 5
	fmt.Println(a)
}
