package main

import "fmt"

func turun(n, i int) {
	if i < 1 {
		return
	}
	fmt.Printf("%d ", i)
	turun(n, i-1)
	if i != 1 {
		fmt.Printf("%d ", i)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)
	turun(n, n)
	fmt.Println()
}