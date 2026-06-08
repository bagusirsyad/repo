package main

import "fmt"

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

var pustaka DaftarBuku
var nPustaka int

func DaftarkanBuku(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&pustaka[i].id)
		fmt.Scan(&pustaka[i].judul)
		fmt.Scan(&pustaka[i].penulis)
		fmt.Scan(&pustaka[i].penerbit)
		fmt.Scan(&pustaka[i].eksemplar)
		fmt.Scan(&pustaka[i].tahun)
		fmt.Scan(&pustaka[i].rating)
	}
}

func CetakTerfavorit(n int) {
	maxIdx := 0
	for i := 1; i < n; i++ {
		if pustaka[i].rating > pustaka[maxIdx].rating {
			maxIdx = i
		}
	}
	b := pustaka[maxIdx]
	fmt.Println(b.judul, b.penulis, b.penerbit, b.tahun, b.rating)
}

func UrutBuku(n int) {
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1
		for j >= 0 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

func Cetak5Terbaru(n int) {
	batas := 5
	if n < 5 {
		batas = n
	}
	for i := 0; i < batas; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

func CariBuku(n, r int) {
	low := 0
	high := n - 1
	hasil := -1
	for low <= high {
		mid := (low + high) / 2
		if pustaka[mid].rating == r {
			hasil = mid
			break
		} else if pustaka[mid].rating < r {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if hasil == -1 {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	} else {
		b := pustaka[hasil]
		fmt.Println(b.judul, b.penulis, b.penerbit, b.tahun, b.eksemplar, b.rating)
	}
}

func main() {
	var n, r int
	fmt.Scan(&n)
	DaftarkanBuku(n)
	CetakTerfavorit(n)
	UrutBuku(n)
	Cetak5Terbaru(n)
	fmt.Scan(&r)
	CariBuku(n, r)
}