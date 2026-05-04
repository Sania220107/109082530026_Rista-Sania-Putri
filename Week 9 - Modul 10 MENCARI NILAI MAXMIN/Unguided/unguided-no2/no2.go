package main

import "fmt"


func hitungTotal(ikan [1000]float64, x, y int) ([1000]float64, int) {
	var total [1000]float64
	jumlahWadah := (x + y - 1) / y

	index := 0
	for i := 0; i < jumlahWadah; i++ {
		count := 0
		for count < y && index < x {
			total[i] += ikan[index]
			index++
			count++
		}
	}
	return total, jumlahWadah
}


func hitungRata(total [1000]float64, jumlahWadah int) float64 {
	var sum float64
	for i := 0; i < jumlahWadah; i++ {
		sum += total[i]
	}
	return sum / float64(jumlahWadah)
}

func main() {
	var x, y int


	fmt.Print("Masukkan jumlah ikan (x) dan kapasitas wadah (y):")
	fmt.Scan(&x, &y)

	var ikan [1000]float64

	fmt.Println("Masukkan berat masing-masing ikan:")
	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	
	total, jumlahWadah := hitungTotal(ikan, x, y)
	rata := hitungRata(total, jumlahWadah)

	fmt.Println("\nTotal berat ikan di setiap wadah")
	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("Wadah ke-%d = %.2f\n", i+1, total[i])
	}

	fmt.Println("\nRata-rata berat ikan per wadah:")
	fmt.Printf("%.2f\n", rata)
}