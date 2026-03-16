package main

import "fmt"

func hitungPersegi(sisi int) {
	var luas, keliling int
	luas = sisi * sisi
	keliling = 4 * sisi

	fmt.Println("Luas Persegi : ", luas)
	fmt.Println("Keliling Persegi : ", keliling)
}

func hitungPersegiPanjang(panjang, lebar int) {
	var luas, keliling int
	
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)

	fmt.Println("Luas persegi panjang : ", luas)
	fmt.Println("Keliling persegi panjang : ", keliling)
}

func hitungLingkaran(jari_jari float64) {
	var phi, luas, keliling float64
	phi = 3.14

	luas = phi * jari_jari*jari_jari
	keliling = 2 * phi * jari_jari

	fmt.Println("Luas lingkaran : ", luas)
	fmt.Println("Keliling lingkaran : ", keliling)
}

func main() {
	var pilihan int

	fmt.Println("--- PROGRAM BANGUN DATAR ---")
	fmt.Println("1. Hitung luas & keliling persegi")
	fmt.Println("2. Hitung luas & keliling persegi panjang")
	fmt.Println("3. Hitung luas & keliling lingkaran")

	fmt.Print("Pilihan : ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		var sisi int
		fmt.Print("Masukan sisi : ")
		fmt.Scan(&sisi)
		hitungPersegi(sisi)
	case 2:
		var p, l int

		fmt.Print("Masukan panjang : ")
		fmt.Scan(&p)
		fmt.Print("Masukan lebar : ")
		fmt.Scan(&l)

		hitungPersegiPanjang(p, l)

	case 3:
		var r float64

		fmt.Print("Masukan jari-jari : ")
		fmt.Scan(&r)

		hitungLingkaran(r)
	}
}