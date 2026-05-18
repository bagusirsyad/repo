package main

import "fmt"

func cetakBintang(baris, n int) {
	if baris > n {
		return
	}
	for i := 0; i < baris; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	cetakBintang(baris+1, n)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)
	cetakBintang(1, n)
}