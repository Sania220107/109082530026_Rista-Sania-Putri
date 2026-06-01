package main

import "fmt"

type arrInt [1000]int

func insertionSort(T *arrInt, n int) {
	var i, pass, temp int

	pass = 1
	for pass < n {
		temp = T[pass]
		i = pass - 1

		for i >= 0 && T[i] > temp {
			T[i+1] = T[i]
			i = i - 1
		}

		T[i+1] = temp
		pass = pass + 1
	}
}

func main() {
	var data arrInt
	var n, x, i, jarak int
	var tetap bool

	n = 0

	fmt.Scan(&x)
	for x >= 0 {
		data[n] = x
		n = n + 1
		fmt.Scan(&x)
	}

	insertionSort(&data, n)

	for i = 0; i < n; i++ {
		fmt.Print(data[i])
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	if n > 1 {
		jarak = data[1] - data[0]
		tetap = true

		for i = 2; i < n; i++ {
			if data[i]-data[i-1] != jarak {
				tetap = false
			}
		}

		if tetap {
			fmt.Println("Data berjarak", jarak)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	}
}