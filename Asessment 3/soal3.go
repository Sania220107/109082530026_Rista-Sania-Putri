package main

import "fmt"

const NMAX = 1000000

type partai struct {
	nama  int
	suara int
}

type tabPartai [NMAX]partai

func main() {
	var p tabPartai
	var n int
	var pilih int
	var idx int
	var temp partai
	var pass, i int

	n = 0

	fmt.Println("Masukkan proses input suara : ")
	fmt.Scan(&pilih)

	for pilih != -1 {
		idx = posisi(p, n, pilih)

		if idx == -1 {
			p[n].nama = pilih
			p[n].suara = 1
			n++
		} else {
			p[idx].suara++
		}

		fmt.Scan(&pilih)
	}

	for pass = 1; pass < n; pass++ {
		temp = p[pass]
		i = pass - 1

		for i >= 0 && p[i].suara < temp.suara {
			p[i+1] = p[i]
			i--
		}

		p[i+1] = temp
	}


	fmt.Println("\nHasil perhitungan suara : ")
	for i = 0; i < n; i++ {
		fmt.Printf("%d(%d)", p[i].nama, p[i].suara)

		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func posisi(t tabPartai, n int, nama int) int {

	var i int

	for i = 0; i < n; i++ {
		if t[i].nama == nama {
			return i
		}
	}

	return -1
}