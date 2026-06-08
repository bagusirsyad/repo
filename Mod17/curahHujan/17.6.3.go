package main

import (
	"fmt"
	"math/rand"
)

func daerah(x, y float64) string {
	if x < 0.5 && y < 0.5 {
		return "A"
	} else if x >= 0.5 && y < 0.5 {
		return "B"
	} else if x >= 0.5 && y >= 0.5 {
		return "C"
	} else {
		return "D"
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	var a, b, c, d int

	for i := 0; i < n; i++ {
		x := rand.Float64()
		y := rand.Float64()
		switch daerah(x, y) {
		case "A":
			a++
		case "B":
			b++
		case "C":
			c++
		case "D":
			d++
		}
	}

	fmt.Printf("Curah hujan daerah A: %.4f milimeter\n", float64(a)*0.0001)
	fmt.Printf("Curah hujan daerah B: %.4f milimeter\n", float64(b)*0.0001)
	fmt.Printf("Curah hujan daerah C: %.4f milimeter\n", float64(c)*0.0001)
	fmt.Printf("Curah hujan daerah D: %.4f milimeter\n", float64(d)*0.0001)
}