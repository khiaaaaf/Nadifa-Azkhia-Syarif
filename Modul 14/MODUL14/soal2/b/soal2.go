package main

import "fmt"

const NMAX = 100

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [NMAX]Buku

func DaftarkanBuku(P *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(
			&P[i].id,
			&P[i].judul,
			&P[i].penulis,
			&P[i].penerbit,
			&P[i].eksemplar,
			&P[i].tahun,
			&P[i].rating,
		)
	}
}

func CetakTerfavorit(P DaftarBuku, n int) {
	max := P[0]

	for i := 1; i < n; i++ {
		if P[i].rating > max.rating {
			max = P[i]
		}
	}

	fmt.Println("Buku Terfavorit:")
	fmt.Println(max.judul, max.penulis, max.penerbit, max.tahun)
}

func UrutBuku(P *DaftarBuku, n int) {
	var temp Buku
	var j int

	for i := 1; i < n; i++ {
		temp = P[i]
		j = i

		for j > 0 && temp.rating > P[j-1].rating {
			P[j] = P[j-1]
			j--
		}

		P[j] = temp
	}
}

func Cetak5Terbaru(P DaftarBuku, n int) {
	fmt.Println("5 Buku Rating Tertinggi:")

	if n > 5 {
		n = 5
	}

	for i := 0; i < n; i++ {
		fmt.Println(P[i].judul)
	}
}

func CariBuku(P DaftarBuku, n, r int) {
	low := 0
	high := n - 1
	found := false

	for low <= high {
		mid := (low + high) / 2

		if P[mid].rating == r {
			found = true
			fmt.Println("Data Buku:")
			fmt.Println(P[mid].judul)
			fmt.Println(P[mid].penulis)
			fmt.Println(P[mid].penerbit)
			fmt.Println(P[mid].tahun)
			fmt.Println(P[mid].eksemplar)
			fmt.Println(P[mid].rating)
			break
		} else if P[mid].rating < r {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if !found {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var P DaftarBuku
	var n, ratingCari int

	fmt.Scan(&n)

	DaftarkanBuku(&P, n)

	CetakTerfavorit(P, n)

	UrutBuku(&P, n)

	Cetak5Terbaru(P, n)

	fmt.Scan(&ratingCari)

	CariBuku(P, n, ratingCari)
}
