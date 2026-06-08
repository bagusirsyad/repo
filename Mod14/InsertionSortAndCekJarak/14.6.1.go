package main

import "fmt"

func insertionSort(arr []int, n int) {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	var data [1000]int
	n := 0

	for {
		var x int
		fmt.Scan(&x)
		if x < 0 {
			break
		}
		data[n] = x
		n++
	}

	insertionSort(data[:], n)

	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(data[i])
	}
	fmt.Println()

	jarak := data[1] - data[0]
	tetap := true
	for i := 2; i < n; i++ {
		if data[i]-data[i-1] != jarak {
			tetap = false
			break
		}
	}

	if tetap {
		fmt.Println("Data berjarak", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}