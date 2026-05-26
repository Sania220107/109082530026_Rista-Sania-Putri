package main

import "fmt"

func HitungSuara() {
	var suara [21]int 
	var input int
	var suaraMasuk int = 0
	var suaraSah int = 0

	for {
		fmt.Scan(&input)


		if input == 0 {
			break
		}
		suaraMasuk++


		if input >= 1 && input <= 20 {
			suara[input]++
			suaraSah++
		}
	}

	fmt.Printf("Suara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)


	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
	}
}

func main() {
	HitungSuara()
}