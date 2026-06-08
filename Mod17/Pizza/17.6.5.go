package main

import "fmt"

func randFloat(seed *int64) float64 {
	*seed = (*seed*1103515245 + 12345) & 0x7fffffff
	return float64(*seed) / float64(0x7fffffff)
}

func dalamLingkaran(x, y, xc, yc, r float64) bool {
	return (x-xc)*(x-xc)+(y-yc)*(y-yc) <= r*r
}

func main() {
	var n int
	fmt.Print("Banyak Topping: ")
	fmt.Scan(&n)

	var seed int64 = 42
	xc, yc, r := 0.5, 0.5, 0.5
	count := 0

	for i := 0; i < n; i++ {
		x := randFloat(&seed)
		y := randFloat(&seed)
		if dalamLingkaran(x, y, xc, yc, r) {
			count++
		}
	}

	fmt.Printf("Topping pada Pizza: %d\n", count)

	pi := 4.0 * float64(count) / float64(n)
	fmt.Printf("PI : %.10f\n", pi)
}