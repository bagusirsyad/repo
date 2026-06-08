package main

import "fmt"

func main() {
	var x, y int
	var berat [1000]float64

	fmt.Scan(&x, &y)
	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	jumlahWadah := x / y
	if x%y != 0 {
		jumlahWadah++
	}

	for w := 0; w < jumlahWadah; w++ {
		awal := w * y
		akhir := awal + y
		if akhir > x {
			akhir = x
		}

		total := 0.0
		for i := awal; i < akhir; i++ {
			total += berat[i]
		}

		if w > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%.2f", total)
	}
	fmt.Println()

	for w := 0; w < jumlahWadah; w++ {
		awal := w * y
		akhir := awal + y
		if akhir > x {
			akhir = x
		}

		total := 0.0
		jumlah := float64(akhir - awal)
		for i := awal; i < akhir; i++ {
			total += berat[i]
		}

		if w > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%.2f", total/jumlah)
	}
	fmt.Println()
}