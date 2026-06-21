package main

import "fmt"

type Domino struct {
	Sisi1 int
	Sisi2 int
}

func gambarKartu(d Domino, suit int) int {

	if suit == 1 {
		return d.Sisi1
	}

	return d.Sisi2
}

func main() {

	kartu := Domino{3, 6}

	fmt.Println(gambarKartu(kartu, 1))
	fmt.Println(gambarKartu(kartu, 2))
}
