package main

import "fmt"

func main() {
	var angka int
	fmt.Print("Masukan angka: ")
	fmt.Scan(&angka)

	for angka < 10 {
		fmt.Println("angka sekarang: ", angka)
		angka++
	}
}