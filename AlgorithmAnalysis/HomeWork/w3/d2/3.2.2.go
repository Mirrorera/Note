package d2

import (
	"fmt"
)

func t2() {
	var n int
	fmt.Scanln(&n)
	n++
	var N string
	N = fmt.Sprintf("%d", n)

	var oper []int
	for _, x := range N {
		oper = append(oper, int(x-'0'))
	}
	//fmt.Println(oper)
	var flag int = -1
	L := len(oper)
	for i, _ := range oper {
		if i == L-1 {
			break
		}
		if oper[i] == oper[i+1] {
			flag = i + 1
			break
		}
	}
	fmt.Print("ANS = ")
	if flag < 0 {
		fmt.Println(n)
		return
	}
	for flag >= 0 {
		oper[flag]++
		if flag == 0 {
			if oper[flag] > 9 {
				fmt.Print("1")
				oper[flag] = 0
				flag = -1
			}
			break
		}
		if oper[flag] > 9 {
			flag--
			continue
		}
		if oper[flag] == oper[flag-1] {
			flag--
			continue
		}
		break
	}

	for i := 0; i <= flag; i++ {
		fmt.Print(oper[i])
	}
	flag++
	for i := 0; i+flag < L; i++ {
		fmt.Print(i % 2)
	}
	fmt.Print("\n")
}
