package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune


func isiArray(t *tabel, n *int) {
	var ch string
	*n = 0

	fmt.Print("Hves : ")
	fmt.Scan(&ch)

	for ch != "." && *n < NMAX {
		t[*n] = rune(ch[0])
		*n++

		fmt.Scan(&ch)
	}
}


func cetakArray(t tabel, n int) {
	fmt.Print("Rvvuusv tves : ")
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}


func balikanArray(t *tabel, n int) {
	var i int
	var temp rune

	for i = 0; i < n/2; i++ {
		temp = t[i]
		t[i] = t[n-1-i]
		t[n-1-i] = temp
	}
}


func palindrom(t tabel, n int) bool {
	var salin tabel
	var i int

	
	for i = 0; i < n; i++ {
		salin[i] = t[i]
	}

	
	balikanArray(&salin, n)

	
	for i = 0; i < n; i++ {
		if t[i] != salin[i] {
			return false
		}
	}

	return true
}

func main() {
	var tab tabel
	var m int

	
	isiArray(&tab, &m)

	
	if palindrom(tab, m) {
		fmt.Println("flalitauou : true")
	} else {
		fmt.Println("flalitauou : false")
	}

	
	balikanArray(&tab, m)

	cetakArray(tab, m)
}