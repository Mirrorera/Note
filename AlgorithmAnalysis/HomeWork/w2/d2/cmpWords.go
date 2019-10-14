package d2

import "fmt"

var s1, s2 string
var count []int

func cmpWords() {

	fmt.Scanln(&s1)
	fmt.Scanln(&s2)

	count := make([]int, 26)
	for _, x := range s1 {
		count[int(x-'a')]++
	}
	for _, x := range s2 {
		count[int(x-'a')]--
	}
	for _, x := range count {
		if x != 0 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
