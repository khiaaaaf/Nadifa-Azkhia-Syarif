package main

import "fmt"

type Domino struct {
	Nilai int
}

func sepasangKartu(a Domino, b Domino) bool {

	return a.Nilai+b.Nilai == 12
}

func main() {

	k1 := Domino{5}
	k2 := Domino{7}

	fmt.Println(sepasangKartu(k1, k2))
}
