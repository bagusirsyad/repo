package main
import "fmt"
func factorial(n int) int64 {
	var hasil int64 = 1
	for i := 1; i <= n; i++ {
		hasil *= int64(i)
	}
	return hasil
}

func permutation(n, r int) int64 {
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int64 {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)

	fmt.Println(permutation(a, c), combination(a, c))
	fmt.Println(permutation(b, d), combination(b, d))
}
