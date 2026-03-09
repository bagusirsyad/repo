package main

import "fmt"

func main() {

	var bil int
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&bil)

	if bil <= 1 {
		fmt.Println("Bilangan harus lebih dari 1")
		return
	}

	fmt.Print("Faktor: ")

	jumlahFaktor := 0

	for i := 1; i <= bil; i++ {
		if bil%i == 0 {
			fmt.Print(i, " ")
			jumlahFaktor++
		}
	}

	fmt.Println()

	if jumlahFaktor == 2 {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}