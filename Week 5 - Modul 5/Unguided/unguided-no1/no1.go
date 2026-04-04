package main

import "fmt"

func polaBintang(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	var n int

	fmt.Scan(&n)
	polaBintang(n)
}