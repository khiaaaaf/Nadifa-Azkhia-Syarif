package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Domino struct {
	Sisi1 int
	Sisi2 int
}

type Dominoes struct {
	Kartu  [28]Domino
	Jumlah int
}

func kocokKartu(d *Dominoes) {
	rand.Seed(time.Now().UnixNano())

	for i := range d.Kartu {
		j := rand.Intn(len(d.Kartu))
		d.Kartu[i], d.Kartu[j] = d.Kartu[j], d.Kartu[i]
	}
}

func main() {
	var deck Dominoes

	kocokKartu(&deck)

	fmt.Println("Kartu berhasil dikocok")
}
