package main

import (
	"fmt"
	"math/rand"
)

type Domino struct {
	suit    [2]int
	nilai   int
	isBalak bool
}

type Dominoes struct {
	kartu        [28]Domino
	jumlahTersisa int
}

func buatDomino(s1, s2 int) Domino {
	return Domino{
		suit:    [2]int{s1, s2},
		nilai:   s1 + s2,
		isBalak: s1 == s2,
	}
}

func buatDominoes() Dominoes {
	var d Dominoes
	idx := 0
	for i := 0; i <= 6; i++ {
		for j := i; j <= 6; j++ {
			d.kartu[idx] = buatDomino(i, j)
			idx++
		}
	}
	d.jumlahTersisa = 28
	return d
}

func kocokKartu(d *Dominoes) {
	for i := 27; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.kartu[i], d.kartu[j] = d.kartu[j], d.kartu[i]
	}
}

func ambilKartu(d *Dominoes) Domino {
	if d.jumlahTersisa == 0 {
		fmt.Println("Tidak ada kartu tersisa!")
		return Domino{}
	}
	kartu := d.kartu[d.jumlahTersisa-1]
	d.jumlahTersisa--
	return kartu
}

func gambarKartu(d Domino, suit int) int {
	if suit == 0 {
		return d.suit[0]
	}
	return d.suit[1]
}

func nilaiKartu(d Domino) int {
	return d.nilai
}

func cetakKartu(d Domino) {
	balak := "bukan balak"
	if d.isBalak {
		balak = "balak"
	}
	fmt.Printf("[%d|%d] nilai=%d (%s)\n", d.suit[0], d.suit[1], d.nilai, balak)
}

func main() {
	set := buatDominoes()
	fmt.Println("=== Set Domino Sebelum Dikocok ===")
	for i := 0; i < 5; i++ {
		cetakKartu(set.kartu[i])
	}

	kocokKartu(&set)
	fmt.Println("\n=== Set Domino Setelah Dikocok (5 kartu pertama) ===")
	for i := 0; i < 5; i++ {
		cetakKartu(set.kartu[i])
	}

	fmt.Printf("\nJumlah kartu tersisa: %d\n", set.jumlahTersisa)

	kartu1 := ambilKartu(&set)
	fmt.Println("\n=== Kartu yang Diambil ===")
	cetakKartu(kartu1)
	fmt.Printf("Jumlah kartu tersisa: %d\n", set.jumlahTersisa)

	fmt.Printf("\nGambar sisi kiri kartu: %d\n", gambarKartu(kartu1, 0))
	fmt.Printf("Gambar sisi kanan kartu: %d\n", gambarKartu(kartu1, 1))
	fmt.Printf("Nilai kartu: %d\n", nilaiKartu(kartu1))
}