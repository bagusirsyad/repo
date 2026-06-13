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
	kartu         [28]Domino
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

func galiKartu(d *Dominoes, target Domino) Domino {
	var hasil Domino
	for d.jumlahTersisa > 0 {
		hasil = ambilKartu(d)
		if hasil.suit[0] == target.suit[0] || hasil.suit[0] == target.suit[1] ||
			hasil.suit[1] == target.suit[0] || hasil.suit[1] == target.suit[1] {
			return hasil
		}
	}
	return hasil
}

func sepasangKartu(a, b Domino) bool {
	return nilaiKartu(a)+nilaiKartu(b) == 12
}

func main() {
	set := buatDominoes()
	kocokKartu(&set)

	fmt.Println("=== Set Domino Setelah Dikocok (5 kartu pertama) ===")
	for i := 0; i < 5; i++ {
		cetakKartu(set.kartu[set.jumlahTersisa-1-i])
	}

	fmt.Println("\n=== galiKartu ===")
	target := ambilKartu(&set)
	fmt.Print("Kartu target: ")
	cetakKartu(target)
	fmt.Printf("Jumlah kartu tersisa setelah ambil target: %d\n", set.jumlahTersisa)

	hasil := galiKartu(&set, target)
	fmt.Print("Kartu hasil gali (suit sama dengan target): ")
	cetakKartu(hasil)
	fmt.Printf("Jumlah kartu tersisa setelah gali: %d\n", set.jumlahTersisa)

	fmt.Println("\n=== sepasangKartu ===")
	kartuA := buatDomino(5, 2)
	kartuB := buatDomino(3, 2)
	fmt.Print("Kartu A: ")
	cetakKartu(kartuA)
	fmt.Print("Kartu B: ")
	cetakKartu(kartuB)
	fmt.Printf("Total nilai: %d\n", nilaiKartu(kartuA)+nilaiKartu(kartuB))
	fmt.Printf("Sepasang (total=12): %v\n", sepasangKartu(kartuA, kartuB))

	kartuC := buatDomino(6, 6)
	kartuD := buatDomino(0, 0)
	fmt.Print("\nKartu C: ")
	cetakKartu(kartuC)
	fmt.Print("Kartu D: ")
	cetakKartu(kartuD)
	fmt.Printf("Total nilai: %d\n", nilaiKartu(kartuC)+nilaiKartu(kartuD))
	fmt.Printf("Sepasang (total=12): %v\n", sepasangKartu(kartuC, kartuD))
}