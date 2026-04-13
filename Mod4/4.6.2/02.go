package main
import "fmt"
func hitungSkor(soal, skor *int) {
	var waktu int

	*soal = 0
	*skor = 0

	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)
		if waktu < 301 {
			*soal = *soal + 1
			*skor += waktu
		}
	}
}

func main() {
	var nama string
	var pemenang string
	var soal, skor int
	var maxSoal, minSkor int
	var pertama bool = true

	for {
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}

		hitungSkor(&soal, &skor)

		if pertama || soal > maxSoal || (soal == maxSoal && skor < minSkor) {
			pemenang = nama
			maxSoal = soal
			minSkor = skor
			pertama = false
		}
	}

	fmt.Println(pemenang, maxSoal, minSkor)
}
