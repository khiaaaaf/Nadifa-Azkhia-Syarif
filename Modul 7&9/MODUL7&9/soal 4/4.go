package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var input string
	fmt.Scan(&input)

	*n = 0
	for i := 0; i < len(input) && input[i] != '.' && *n < NMAX; i++ {
		t[*n] = rune(input[i])
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks : ")
	isiArray(&tab, &m)

	balikanArray(&tab, m)

	fmt.Print("Reverse teks : ")
	cetakArray(tab, m)
}
