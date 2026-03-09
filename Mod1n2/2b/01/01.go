package main

import "fmt"

func main() {

	pola := []string{"merah", "kuning", "hijau", "ungu"}
	hasil := true

	for i := 1; i <= 5; i++ {

		for j := 0; j < 4; j++ {

			var warna string
			fmt.Scan(&warna)

			if warna != pola[j] {
				hasil = false
			}
		}
	}

	fmt.Println(hasil)
}