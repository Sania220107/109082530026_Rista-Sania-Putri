package main

import "fmt"

const nMax int = 51

type mhs struct {
	NIM   string
	nama  string
	nilai int
}

type arrMahasiswa [nMax]mhs

func main() {
	var data arrMahasiswa
	var jumlahData int
	var index1, index2 int
	var NIM string


	inputMhs(&data, &jumlahData)


	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&NIM)

	index1 = findFirstScore(data, jumlahData, NIM)
	index2 = findMaxScore(data, jumlahData, NIM)

	if index1 != -1 {
		fmt.Println("Nilai pertama dari NIM", NIM, "adalah", data[index1].nilai)
		fmt.Println("Nilai terbesar dari NIM", NIM, "adalah", data[index2].nilai)
	} else {
		fmt.Println("NIM", NIM, "tidak ditemukan")
	}
}

func inputMhs(T *arrMahasiswa, N *int) {
	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(N)

	for i := 0; i < *N; i++ {
		fmt.Print("Masukkan data ke-", i+1, " : ")
		fmt.Scan(&T[i].NIM, &T[i].nama, &T[i].nilai)
	}
}

func findFirstScore(T arrMahasiswa, N int, nim string) int {
	found := -1
	i := 0

	for i < N && found == -1 {
		if T[i].NIM == nim {
			found = i
		}
		i++
	}

	return found
}

func findMaxScore(T arrMahasiswa, N int, nim string) int {
	found := findFirstScore(T, N, nim)

	if found == -1 {
		return -1
	}

	indexMax := found

	for i := found; i < N; i++ {
		if T[i].NIM == nim && T[i].nilai > T[indexMax].nilai {
			indexMax = i
		}
	}

	return indexMax
}