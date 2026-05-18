package main

import "fmt"

func cetakFaktor(n, pembagi int) {
	if pembagi > n {
		return
	}
	if n%pembagi == 0 {
		fmt.Printf("%d ", pembagi)
	}
	cetakFaktor(n, pembagi+1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)
	fmt.Printf("Faktor dari %d: ", n)
	cetakFaktor(n, 1)
	fmt.Println()
}