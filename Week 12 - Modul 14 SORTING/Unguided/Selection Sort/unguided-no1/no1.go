package main

import "fmt"

type arrRumah [1000000]int
type matriksRumah [1000]arrRumah
type arrJumlah [1000]int

func selectionSort(T *arrRumah, n int) {
	var i, j, idxMin, temp int

	i = 1
	for i <= n-1 {
		idxMin = i - 1
		j = i

		for j < n {
			if T[idxMin] > T[j] {
				idxMin = j
			}
			j = j + 1
		}

		temp = T[idxMin]
		T[idxMin] = T[i-1]
		T[i-1] = temp

		i = i + 1
	}
}

func main() {
	var data matriksRumah
	var jumlah arrJumlah
	var n, m, i, j int

	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&m)
		jumlah[i] = m

		for j = 0; j < m; j++ {
			fmt.Scan(&data[i][j])
		}

		selectionSort(&data[i], m)
	}

	for i = 0; i < n; i++ {
		for j = 0; j < jumlah[i]; j++ {
			fmt.Print(data[i][j])

			if j < jumlah[i]-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}