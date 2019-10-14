package d2

import "fmt"

func qpow(n float64, k int) float64 {
	var ans float64 = 1
	for k > 0 {
		if k%2 == 1 {
			ans = ans * n
		}
		n = n * n
		k = k >> 1
	}
	return ans
}
func d2() {
	var N int
	fmt.Scanln(&N)
	var ans float64

	ans = (2.0 / 9.0) * ((qpow(10.0, (N+1))-10)/9.0 - float64(N))

	fmt.Println("ANS =", ans)
}
