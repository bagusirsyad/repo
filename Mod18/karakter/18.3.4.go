package main

import "fmt"

var (
	input []rune
	pos   int
)

func start(s string) {
	input = []rune(s)
	pos = 0
}

func maju() {
	if pos < len(input) {
		pos++
	}
}

func eop() bool {
	return pos >= len(input) || input[pos] == '.'
}

func cc() rune {
	return input[pos]
}

func bacaSemua() {
	fmt.Print("Karakter terbaca: ")
	for !eop() {
		fmt.Printf("%c", cc())
		maju()
	}
	fmt.Println()
}

func hitungKarakter() int {
	jumlah := 0
	for !eop() {
		jumlah++
		maju()
	}
	return jumlah
}

func hitungHurufA() int {
	jumlah := 0
	for !eop() {
		if cc() == 'A' {
			jumlah++
		}
		maju()
	}
	return jumlah
}

func hitungFrekuensiA() float64 {
	jumlahA := 0
	total := 0
	for !eop() {
		total++
		if cc() == 'A' {
			jumlahA++
		}
		maju()
	}
	if total == 0 {
		return 0
	}
	return float64(jumlahA) / float64(total)
}

func hitungKataLE() int {
	jumlah := 0
	prev := rune(0)
	for !eop() {
		if prev == 'L' && cc() == 'E' {
			jumlah++
		}
		prev = cc()
		maju()
	}
	return jumlah
}

func main() {
	teks := "HALO SELAMAT DATANG KE TELKOM UNIVERSITY."

	fmt.Println("Input:", teks)
	fmt.Println()

	start(teks)
	bacaSemua()

	start(teks)
	fmt.Printf("Jumlah karakter: %d\n", hitungKarakter())

	start(teks)
	fmt.Printf("Jumlah huruf A: %d\n", hitungHurufA())

	start(teks)
	fmt.Printf("Frekuensi huruf A: %.4f\n", hitungFrekuensiA())

	start(teks)
	fmt.Printf("Jumlah kata LE: %d\n", hitungKataLE())
}
