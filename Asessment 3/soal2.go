package main

import "fmt"

const NMAX = 1001

type Pemain struct {
	namaDepan    string
	namaBelakang string
	gol          int
	assist       int
}

type arrPemain [NMAX]Pemain

func selectionSort(A *arrPemain, n int) {
	var pass, idxMax, i int
	var temp Pemain

	for pass = 0; pass < n-1; pass++ {
		idxMax = pass

		for i = pass + 1; i < n; i++ {

			
			if A[i].gol > A[idxMax].gol {
				idxMax = i

			} else if A[i].gol == A[idxMax].gol {
				if A[i].assist > A[idxMax].assist {
					idxMax = i
				}
			}
		}

		
		temp = A[pass]
		A[pass] = A[idxMax]
		A[idxMax] = temp
	}
}

func main() {
	var A arrPemain
	var n, i int

	fmt.Println("Masukkan Data Input : ")
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&A[i].namaDepan,
			&A[i].namaBelakang,
			&A[i].gol,
			&A[i].assist)
	}

	selectionSort(&A, n)

	fmt.Println("\nHasil Sorting : ")
	for i = 0; i < n; i++ {
		fmt.Println(
			A[i].namaDepan,
			A[i].namaBelakang,
			A[i].gol,
			A[i].assist,
		)
	}
}