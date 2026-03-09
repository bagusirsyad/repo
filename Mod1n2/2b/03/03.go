package main
import (
	"fmt"
	"math"
)
func main() {

	for {

		var kiri, kanan float64

		fmt.Print("Masukkan berat kantong kiri dan kanan: ")
		fmt.Scan(&kiri, &kanan)

		if kiri < 0 || kanan < 0 || (kiri+kanan) > 150 {
			fmt.Println("Proses selesai")
			break
		}

		selisih := math.Abs(kiri - kanan)

		if selisih >= 9 {
			fmt.Println("Sepeda motor akan oleng: true")
		} else {
			fmt.Println("Sepeda motor akan oleng: false")
		}
	}
}