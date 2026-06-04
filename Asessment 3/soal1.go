package main

import "fmt"

const NMAX = 1000000

type arrInt [NMAX]int

func sorting(T *arrInt, n int) {

	var pass, idxMin, i, temp int

	for pass = 0; pass < n-1; pass++ {
		idxMin = pass

		for i = pass + 1; i < n; i++ {
			if T[i] < T[idxMin] {
				idxMin = i
			}
		}

		temp = T[pass]
		T[pass] = T[idxMin]
		T[idxMin] = temp
	}
}

func median(T arrInt, n int) float64 {

	fmt.Println("Median : ")
	if n%2 == 1 {
		return float64(T[n/2])
	}

	return float64(T[n/2-1]+T[n/2]) / 2.0
}

func main() {
	var T arrInt
	var n int = 0
	var x int

	fmt.Println("Input data masukan : ")
	for {
		
		fmt.Scan(&x)

		if x == -5313541 {
			break
		}

		if x == 0 {
			sorting(&T, n)
			fmt.Println(median(T, n))
		} else {
			T[n] = x
			n++
		}
	}
}