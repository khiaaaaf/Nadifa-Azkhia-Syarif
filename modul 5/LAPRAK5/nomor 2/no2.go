package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	bintang(n, 1)
}

func bintang(n int, i int) {
	if i > n {
		return
	}
	fmt.Println(generateBintang(i))
	bintang(n, i+1)
}

func generateBintang(jumlah int) string {
	result := ""
	for i := 0; i < jumlah; i++ {
		result += "*"
	}
	return result
}
