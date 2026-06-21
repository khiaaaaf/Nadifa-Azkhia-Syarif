package main

import "fmt"

var data string
var idx int

func start(teks string) {
	data = teks
	idx = 0
}

func maju() {
	idx++
}

func eop() bool {
	return data[idx] == '.'
}

func cc() byte {
	return data[idx]
}

func main() {

	start("BELAJAR.")

	for !eop() {
		fmt.Printf("%c ", cc())
		maju()
	}
}
