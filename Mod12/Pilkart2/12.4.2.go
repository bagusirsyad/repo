package main

import "fmt"

func main() {
	var suara [21]int
	totalMasuk := 0
	totalSah := 0

	for {
		var x int
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		totalMasuk++
		if x >= 1 && x <= 20 {
			suara[x]++
			totalSah++
		}
	}

	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", totalSah)

	ketua := 0
	for i := 1; i <= 20; i++ {
		if suara[i] > suara[ketua] {
			ketua = i
		}
	}

	wakil := 0
	for i := 1; i <= 20; i++ {
		if i != ketua && suara[i] > suara[wakil] {
			wakil = i
		}
	}

	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}