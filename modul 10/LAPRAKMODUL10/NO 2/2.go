package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	var ikan [1000]float64

	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	var totalWadah []float64

	for i := 0; i < x; i += y {
		sum := 0.0
		for j := i; j < i+y && j < x; j++ {
			sum += ikan[j]
		}
		totalWadah = append(totalWadah, sum)
	}

	totalSemua := 0.0
	for i := 0; i < len(totalWadah); i++ {
		fmt.Printf("%.2f ", totalWadah[i])
		totalSemua += totalWadah[i]
	}

	fmt.Println()

	rata := totalSemua / float64(len(totalWadah))
	fmt.Printf("Rata-rata: %.2f\n", rata)
}
