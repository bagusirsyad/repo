package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0
	fmt.Scanf("%c", &ch)
	for ch != '.' && *n < NMAX {
		t[*n] = ch
		*n++
		fmt.Scanf("%c", &ch)
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	var balik tabel
	for i := 0; i < n; i++ {
		balik[i] = t[i]
	}
	balikanArray(&balik, n)
	for i := 0; i < n; i++ {
		if t[i] != balik[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks\t\t: ")
	isiArray(&tab, &m)

	fmt.Print("Reverse teks\t: ")
	balikanArray(&tab, m)
	cetakArray(tab, m)

	balikanArray(&tab, m)

	fmt.Printf("Palindrom\t? %v\n", palindrom(tab, m))
}
