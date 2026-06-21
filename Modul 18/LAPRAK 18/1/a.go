package main

import "fmt"

type Domino struct {
	Sisi1 int
	Sisi2 int
	Nilai int
	Balak bool
}

type Dominoes struct {
	Kartu  [28]Domino
	Jumlah int
}

func main() {
	var kartu Domino

	kartu.Sisi1 = 3
	kartu.Sisi2 = 3
	kartu.Nilai = 6
	kartu.Balak = true

	fmt.Println(kartu)
}
