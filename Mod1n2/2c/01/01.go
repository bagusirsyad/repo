package main

import "fmt"

func main() {

	var beratGram int

	fmt.Print("Masukkan berat parsel (gram): ")
	fmt.Scan(&beratGram)

	kg := beratGram / 1000
	sisaGram := beratGram % 1000
	biayaKg := kg * 10000
	biayaSisa := 0

	if kg > 10 {
		biayaSisa = 0
	} else {
		if sisaGram >= 500 {
			biayaSisa = sisaGram * 5
		} else {
			biayaSisa = sisaGram * 15
		}
	}

	totalBiaya := biayaKg + biayaSisa
	fmt.Println("\n===== DETAIL PENGIRIMAN =====")
	fmt.Println("Berat parsel       :", beratGram, "gram")
	fmt.Println("Detail berat       :", kg, "kg +", sisaGram, "gram")
	fmt.Println("Biaya per kg       : Rp", biayaKg)
	fmt.Println("Biaya sisa gram    : Rp", biayaSisa)
	fmt.Println("Detail biaya       : Rp", biayaKg, "+ Rp", biayaSisa)
	fmt.Println("Total biaya kirim  : Rp", totalBiaya)
}
