package main

import "fmt"

func cariMinMax(data []float64, n int) (float64, float64) {
	min := data[0]
	max := data[0]

	for i := 1; i < n; i++ {
		if data[i] < min {
			min = data[i]
		}
		if data[i] > max {
			max = data[i]
		}
	}

	return min, max
}

func main() {
	var N int
	var berat [1000]float64

	fmt.Print("Masukkan jumlah anak kelinci yang akan ditimbang: ")
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&berat[i])
	}

	min, max := cariMinMax(berat[:], N)

	fmt.Println("Data Berat Anak Kelinci")
	fmt.Println("---------------------------")
	fmt.Printf("Berat terbesar : %.2f\n", max)
	fmt.Printf("Berat terkecil : %.2f\n", min)
}