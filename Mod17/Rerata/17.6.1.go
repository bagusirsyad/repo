package main

import "fmt"

func main() {
	var angka float64
	var jumlah float64
	var count int

	for {
		fmt.Scan(&angka)
		if angka == 9999 {
			break
		}
		jumlah += angka
		count++
	}

	if count == 0 {
		fmt.Println("Tidak ada bilangan yang dimasukkan")
	} else {
		rerata := jumlah / float64(count)
		fmt.Printf("Rerata: %.2f\n", rerata)
	}
}