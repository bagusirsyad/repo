package main

import "fmt"

func akar(x float64) float64 {
	if x == 0 {
		return 0
	}
	z := x
	for i := 0; i < 1000; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("\na. Keseluruhan isi array:")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	fmt.Println("\nb. Elemen dengan indeks ganjil:")
	for i := 1; i < len(arr); i += 2 {
		fmt.Printf("arr[%d]=%d ", i, arr[i])
	}
	fmt.Println()

	fmt.Println("\nc. Elemen dengan indeks genap:")
	for i := 0; i < len(arr); i += 2 {
		fmt.Printf("arr[%d]=%d ", i, arr[i])
	}
	fmt.Println()

	var x int
	fmt.Print("\nd. Masukkan nilai x (kelipatan indeks): ")
	fmt.Scan(&x)
	fmt.Printf("Elemen dengan indeks kelipatan %d:\n", x)
	for i := x; i < len(arr); i += x {
		fmt.Printf("arr[%d]=%d ", i, arr[i])
	}
	fmt.Println()

	var hapus int
	fmt.Print("\ne. Masukkan indeks yang akan dihapus: ")
	fmt.Scan(&hapus)
	arrHapus := append(arr[:hapus], arr[hapus+1:]...)
	fmt.Println("Array setelah penghapusan:")
	for i := 0; i < len(arrHapus); i++ {
		fmt.Printf("%d ", arrHapus[i])
	}
	fmt.Println()

	total := 0
	for i := 0; i < len(arr); i++ {
		total += arr[i]
	}
	rataRata := float64(total) / float64(len(arr))
	fmt.Printf("\nf. Rata-rata: %.2f\n", rataRata)

	jumlah := 0.0
	for i := 0; i < len(arr); i++ {
		jumlah += (float64(arr[i]) - rataRata) * (float64(arr[i]) - rataRata)
	}
	stdDev := akar(jumlah / float64(len(arr)))
	fmt.Printf("\ng. Standar deviasi: %.2f\n", stdDev)

	var cari int
	fmt.Print("\nh. Masukkan bilangan yang dicari frekuensinya: ")
	fmt.Scan(&cari)
	frekuensi := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == cari {
			frekuensi++
		}
	}
	fmt.Printf("Frekuensi %d dalam array: %d\n", cari, frekuensi)
}