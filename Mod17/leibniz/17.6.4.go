package main

import "fmt"

func suku(i int) float64 {
	if i%2 == 1 {
		return 1.0 / float64(2*i-1)
	}
	return -1.0 / float64(2*i-1)
}

func main() {
	var n int
	fmt.Print("N suku pertama: ")
	fmt.Scan(&n)

	var s float64
	for i := 1; i <= n; i++ {
		s += suku(i)
	}

	fmt.Printf("Hasil PI: %.7f\n", s*4)

	var si, siBerikut float64
	si = suku(1)
	var idx int
	for i := 2; ; i++ {
		siBerikut = si + suku(i)
		nilaiSuku := suku(i)
		if nilaiSuku < 0 {
			nilaiSuku = -nilaiSuku
		}
		if nilaiSuku <= 0.00001 {
			idx = i
			break
		}
		si = siBerikut
	}

	fmt.Printf("Hasil PI: %.10f\n", si*4)
	fmt.Printf("Hasil PI: %.10f\n", siBerikut*4)
	fmt.Printf("Pada i ke: %d\n", idx)
}