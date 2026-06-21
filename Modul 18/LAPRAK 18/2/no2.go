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

	d.Jumlah--
	return d.Kartu[d.Jumlah]
}

func galiKartu(deck *Dominoes, target Domino) {

	for deck.Jumlah > 0 {

		kartu := ambilKartu(deck)

		if kartu.Sisi1 == target.Sisi1 ||
			kartu.Sisi2 == target.Sisi1 ||
			kartu.Sisi1 == target.Sisi2 ||
			kartu.Sisi2 == target.Sisi2 {

			fmt.Println("Kartu ditemukan")
			fmt.Println(kartu)
			return
		}
	}

	fmt.Println("Tidak ditemukan")
}

func main() {

	var deck Dominoes

	deck.Kartu[0] = Domino{1, 4}
	deck.Kartu[1] = Domino{5, 6}
	deck.Jumlah = 2

	target := Domino{4, 2}

	galiKartu(&deck, target)
}
