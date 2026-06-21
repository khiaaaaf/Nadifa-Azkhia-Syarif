package main

import "fmt"

type Domino struct {
	Sisi1 int
	Sisi2 int
}

type Dominoes struct {
	Kartu  [28]Domino
	Jumlah int
}

func ambilKartu(d *Dominoes) Domino {

	if d.Jumlah == 0 {
		return Domino{}
	}

	d.Jumlah--
	return d.Kartu[d.Jumlah]
}

func main() {

	var deck Dominoes

	deck.Kartu[0] = Domino{2, 5}
	deck.Jumlah = 1

	kartu := ambilKartu(&deck)

	fmt.Println(kartu)
}
