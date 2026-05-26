package main

import "fmt"

const NMAX = 21

var data [NMAX]int

func HitungSuara() (int, int) {
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
			data[input]++
			suaraSah++
		}
	}

	return suaraMasuk, suaraSah
}

func CariKetuaDanWakil() (int, int) {
	var ketua int = 1
	var wakil int = 2


	for i := 2; i <= 20; i++ {
		if data[i] > data[ketua] {
			ketua = i
		}
	}


	for i := 1; i <= 20; i++ {
		if i != ketua {
			wakil = i
			break
		}
	}

	for i := 1; i <= 20; i++ {
		if i != ketua {
			if data[i] > data[wakil] {
				wakil = i
			}
		}
	}

	return ketua, wakil
}

func main() {
	var suaraMasuk, suaraSah int
	var ketua, wakil int

	suaraMasuk, suaraSah = HitungSuara()
	ketua, wakil = CariKetuaDanWakil()

	fmt.Printf("Suara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)
	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}