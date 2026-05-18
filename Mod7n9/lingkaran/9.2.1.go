package main

import "fmt"

type Titik struct {
	x, y int
}

type Lingkaran struct {
	pusat  Titik
	radius int
}

func jarak(p, q Titik) int {
	return (p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)
}

func didalam(c Lingkaran, p Titik) bool {
	return jarak(c.pusat, p) <= c.radius*c.radius
}

func main() {
	var l1, l2 Lingkaran
	var t Titik

	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.radius)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.radius)
	fmt.Scan(&t.x, &t.y)

	d1 := didalam(l1, t)
	d2 := didalam(l2, t)

	if d1 && d2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if d1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if d2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}