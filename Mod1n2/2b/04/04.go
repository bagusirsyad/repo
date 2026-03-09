package main

import (
	"fmt"
)

func main() {

	var K int
	fmt.Print("Nilai K = ")
	fmt.Scan(&K)

	fk := float64((4*K+2)*(4*K+2)) / float64((4*K+1)*(4*K+3))
	fmt.Printf("Nilai f(K) = %.10f\n", fk)

	total := 0.0

	for k := 0; k <= K; k++ {
		suku := float64((4*k+2)*(4*k+2)) / float64((4*k+1)*(4*k+3))
		total += suku
	}

	fmt.Printf("Nilai akar 2 = %.10f\n", total)
}