package main

import "fmt"

func adaDalam(data []string, x string) bool {
	for i := 0; i < len(data); i++ {
		if data[i] == x {
			return true
		}
	}
	return false
}

func posisiPertama(data []string, x string) int {
	for i := 0; i < len(data); i++ {
		if data[i] == x {
			return i + 1
		}
	}
	return -1
}

func hitungMunculnya(data []string, x string) int {
	count := 0
	for i := 0; i < len(data); i++ {
		if data[i] == x {
			count++
		}
	}
	return count
}

func main() {
	var x string
	var n int

	fmt.Scan(&x)
	fmt.Scan(&n)

	data := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	if adaDalam(data, x) {
		fmt.Println("a. String", x, "ADA dalam kumpulan")
	} else {
		fmt.Println("a. String", x, "TIDAK ADA dalam kumpulan")
	}

	posisi := posisiPertama(data, x)
	if posisi != -1 {
		fmt.Println("b. String", x, "pertama ditemukan pada posisi ke-", posisi)
	} else {
		fmt.Println("b. String", x, "tidak ditemukan")
	}

	jumlah := hitungMunculnya(data, x)
	fmt.Println("c. String", x, "muncul sebanyak", jumlah, "kali")

	if jumlah >= 2 {
		fmt.Println("d. Ada sedikitnya dua string", x, "dalam kumpulan")
	} else {
		fmt.Println("d. Tidak ada sedikitnya dua string", x, "dalam kumpulan")
	}
}