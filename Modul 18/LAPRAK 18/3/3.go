package main

import "fmt"

type Domino struct {
	Sisi1 int
	Sisi2 int
}

func bisaSambung(a Domino, b Domino) bool {

	return a.Sisi1 == b.Sisi1 ||
		a.Sisi1 == b.Sisi2 ||
		a.Sisi2 == b.Sisi1 ||
		a.Sisi2 == b.Sisi2
}

func main() {

	kartu1 := Domino{2, 4}
	kartu2 := Domino{4, 6}

	if bisaSambung(kartu1, kartu2) {
		fmt.Println("Kartu dapat disambung")
	} else {
		fmt.Println("Kartu tidak dapat disambung")
	}
}
