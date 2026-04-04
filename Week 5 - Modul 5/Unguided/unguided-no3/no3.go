package main

import "fmt"

func pangkat(x, y int) int {
	var hasil int

	if y == 0 {
		return 1
	} else {
		hasil = x * (pangkat(x, y-1))
		return  hasil
	}
}

func main() {
	var x, y int
	var hasil int
	fmt.Print("Masukan nilai x dan y: ")
	fmt.Scan(&x, &y)
	hasil = pangkat(x, y)
	fmt.Println(hasil)
}