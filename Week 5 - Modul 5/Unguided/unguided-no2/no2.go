package main

import "fmt"

func bilangan(n int) {
	if n == 1 {
		fmt.Print(n, " ")
		return
	} else {
		fmt.Print(n, " ")
		bilangan(n-1)
		fmt.Print(n, " ")
	}
}


func main() {
	var n int
	fmt.Print("Masukan bilangan n: ")
	fmt.Scan(&n)
	bilangan(n)
}