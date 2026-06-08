package main

import "fmt"

func selectionSort(arr []int) {
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

		selectionSort(rumah)

		fmt.Printf("Urutan rumah kerabat di daerah %d: ", d+1)
		for i := 0; i < m; i++ {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(rumah[i])
		}
		fmt.Println()
	}
}