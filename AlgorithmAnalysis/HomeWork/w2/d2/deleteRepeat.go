package main

import "fmt"

func main() {
	var r []int
	var N int

	fmt.Scanln(&N)
	r = make([]int, N)
	for i, _ := range r {
		fmt.Scanf("%d", &r[i])
	}

	var hash map[int]bool = make(map[int]bool)

	count := 0
	for _, x := range r {
		_, ok := hash[x]
		if ok {
			continue
		}
		hash[x] = true
		r[count] = x
		count++
	}

	var ans []int = r[0:count]
	fmt.Println(ans)
}
