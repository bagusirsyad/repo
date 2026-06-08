package main

import "fmt"

func selectionSortAsc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func selectionSortDesc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah daerah: ")
	fmt.Scan(&n)

	for d := 0; d < n; d++ {
		var m int
		fmt.Printf("Masukkan jumlah rumah kerabat di daerah %d: ", d+1)
		fmt.Scan(&m)

		rumah := make([]int, m)
		fmt.Printf("Masukkan %d nomor rumah: ", m)
		for i := 0; i < m; i++ {
			fmt.Scan(&rumah[i])
		}

		var ganjil, genap []int
		for _, v := range rumah {
			if v%2 != 0 {
				ganjil = append(ganjil, v)
			} else {
				genap = append(genap, v)
			}
		}

		selectionSortAsc(ganjil)
		selectionSortDesc(genap)

		fmt.Printf("Urutan rumah kerabat di daerah %d: ", d+1)
		hasil := append(ganjil, genap...)
		for i, v := range hasil {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(v)
		}
		fmt.Println()
	}
}