package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y int
}

type lingkaran struct {
	pusat titik
	r     int
}

func jarak(p, q titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

func didalam(c lingkaran, p titik) bool {
	return jarak(c.pusat, p) <= float64(c.r)
}

func main() {
	var L1, L2 lingkaran
	var p titik

	fmt.Scan(&L1.pusat.x, &L1.pusat.y, &L1.r)
	fmt.Scan(&L2.pusat.x, &L2.pusat.y, &L2.r)
	fmt.Scan(&p.x, &p.y)

	d1 := didalam(L1, p)
	d2 := didalam(L2, p)

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
