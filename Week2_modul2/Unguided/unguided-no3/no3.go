package main

import "fmt"

func main() {
	var (
		berat, kg, sisa, biaya_kg, biaya_sisa, total_biaya int
	)

	fmt.Print("Masukkan total berat (gram): ")
	fmt.Scan(&berat)

	kg = berat / 1000
	sisa = berat % 1000
	biaya_kg = kg * 10000

	if kg > 10 {
		biaya_sisa = 0
	} else if sisa >= 500 {
		biaya_sisa = sisa * 5
	} else {
		biaya_sisa = sisa * 15
	}

	total_biaya = biaya_kg + biaya_sisa

	fmt.Println("\n===== Detail Perhitungan =====")
	fmt.Println("Detail berat: ", kg, "+", sisa, "gram")
	fmt.Println("Detail Biaya: Rp. ", biaya_kg, "+ Rp.", biaya_sisa)
	fmt.Println("Total biaya: Rp ", total_biaya)
}