package main

import (
	"fmt"
	"math"
)


func tampilSemua(arr [100]int, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}


func indeksGanjil(arr [100]int, n int) {
	for i := 0; i < n; i++ {
		if i%2 != 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
}


func indeksGenap(arr [100]int, n int) {
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
}


func kelipatanX(arr [100]int, n int, x int) {
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
}


func hapusIndex(arr *[100]int, n *int, idx int) {
	for i := idx; i < *n-1; i++ {
		arr[i] = arr[i+1]
	}
	*n--
}

func rataRata(arr [100]int, n int) float64 {
	total := 0
	for i := 0; i < n; i++ {
		total += arr[i]
	}
	return float64(total) / float64(n)
}

func stdDev(arr [100]int, n int) float64 {
	rata := rataRata(arr, n)
	var jumlah float64

	for i := 0; i < n; i++ {
		selisih := float64(arr[i]) - rata
		jumlah += selisih * selisih
	}

	return math.Sqrt(jumlah / float64(n))
}


func frekuensi(arr [100]int, n int, x int) int {
	count := 0
	for i := 0; i < n; i++ {
		if arr[i] == x {
			count++
		}
	}
	return count
}

func main() {
	var arr [100]int
	var n, x, idx, cari int

	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Print("Masukkan x: ")
	fmt.Scan(&x)

	fmt.Print("Index yang dihapus: ")
	fmt.Scan(&idx)

	fmt.Print("Cari angka: ")
	fmt.Scan(&cari)

	hapusIndex(&arr, &n, idx)

	fmt.Println("\nSemua array:")
	tampilSemua(arr, n)

	fmt.Println("Indeks ganjil:")
	indeksGanjil(arr, n)

	fmt.Println("Indeks genap:")
	indeksGenap(arr, n)

	fmt.Println("Kelipatan x:")
	kelipatanX(arr, n, x)

	fmt.Println("Rata-rata:", rataRata(arr, n))
	fmt.Println("Standar deviasi:", stdDev(arr, n))
	fmt.Println("Frekuensi:", frekuensi(arr, n, cari))
}