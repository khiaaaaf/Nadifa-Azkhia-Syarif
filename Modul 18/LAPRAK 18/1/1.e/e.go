package main

import "fmt"

type Domino struct {
	Sisi1 int
	Sisi2 int
	Nilai int
}

func nilaiKartu(d Domino) int {
	return d.Nilai
}

func main() {

	kartu := Domino{2, 5, 7}

	fmt.Println(nilaiKartu(kartu))
}
