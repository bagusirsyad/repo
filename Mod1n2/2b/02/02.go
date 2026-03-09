package main

import "fmt"

func main() {

	var n int
	fmt.Print("Masukkan jumlah bunga: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Jumlah bunga harus lebih dari 0")
		return
	}

	pita := ""

	for i := 1; i <= n; i++ {

		var nama string
		fmt.Printf("Masukkan bunga ke-%d: ", i)
		fmt.Scan(&nama)

		if i == 1 {
			pita = nama
		} else {
			pita = pita + " - " + nama
		}
	}

	fmt.Println("Isi pita:", pita)
}