package main

import "fmt"


func inputKlub(klubA, klubB *string) {
	fmt.Print("Klub A : ")
	fmt.Scan(klubA)

	fmt.Print("Klub B : ")
	fmt.Scan(klubB)
}


func cekPemenang(skorA, skorB int, klubA, klubB string) string {
	if skorA > skorB {
		return klubA
	} else if skorB > skorA {
		return klubB
	}
	return "Draw"
}


func tampilHasil(data [100]string, jumlah int) {
	for i := 0; i < jumlah; i++ {
		fmt.Printf("Hasil %d : %s\n", i+1, data[i])
	}
	fmt.Println("Pertandingan selesai")
}

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var pemenang [100]string
	var jumlah int
	var pertandingan int = 1


	inputKlub(&klubA, &klubB)


	for {
		fmt.Printf("Pertandingan %d : ", pertandingan)
		fmt.Scan(&skorA, &skorB)


		if skorA < 0 || skorB < 0 {
			break
		}


		pemenang[jumlah] = cekPemenang(skorA, skorB, klubA, klubB)
		jumlah++
		pertandingan++
	}


	tampilHasil(pemenang, jumlah)
}