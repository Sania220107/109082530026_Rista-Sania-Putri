package main

import "fmt"

const nProv int = 10

type NamaProv [nProv]string
type PopProv [nProv]int
type TumbuhProv [nProv]float64

func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	var i int

	for i = 0; i < nProv; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int {
	var idx int = 0
	var i int

	for i = 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[idx] {
			idx = i
		}
	}

	return idx
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	var i int
	var found int = -1

	for i = 0; i < nProv && found == -1; i++ {
		if prov[i] == nama {
			found = i
		}
	}

	return found
}

func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	var i int
	var hasil float64

	fmt.Println()
	fmt.Println("=== Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ===")

	for i = 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			hasil = float64(pop[i]) * (1 + tumbuh[i])
			fmt.Println(prov[i], int(hasil))
		}
	}
}

func main() {
	var TProvinsi NamaProv
	var TPopulasi PopProv
	var TPertumbuhan TumbuhProv

	var cari string
	var idxTercepat int
	var idxProvinsi int

	fmt.Println("=== Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ===")

	InputData(&TProvinsi, &TPopulasi, &TPertumbuhan)

	fmt.Println()

	idxTercepat = ProvinsiTercepat(TPertumbuhan)

	fmt.Println("Provinsi dengan angka pertumbuhan tercepat :", TProvinsi[idxTercepat])

	fmt.Println()

	fmt.Print("Data provinsi yang dicari : ")
	fmt.Scan(&cari)

	idxProvinsi = IndeksProvinsi(TProvinsi, cari)

	if idxProvinsi != -1 {
		fmt.Println(TProvinsi[idxProvinsi])
	} else {
		fmt.Println("Provinsi tidak ditemukan")
	}

	Prediksi(TProvinsi, TPopulasi, TPertumbuhan)
}