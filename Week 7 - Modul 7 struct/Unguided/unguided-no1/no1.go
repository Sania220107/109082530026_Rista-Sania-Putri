package main

import "fmt"

type suhu float64

func CelciusToReamur(Celcius suhu) suhu {
	hasil := 4.0 / 5.0 * Celcius
	return hasil
}

func CelciusToFahrenheit(Celcius suhu) suhu {
	hasil := (9.0 / 5.0 * Celcius) + 32
	return hasil
}

func CelciusToKelvin(Celcius suhu) suhu {
	hasil := Celcius + 273.15
	return hasil
}

func main() {
	var celcius suhu

	fmt.Println("=== KONVERTER CELCIUS ===")
	fmt.Print("Masukkan suhu (celcius) : ")
	fmt.Scan(&celcius)

	fmt.Printf("%.0f celcius = %.1f reamur", celcius, CelciusToReamur(suhu(celcius)))
	fmt.Printf("\n%.0f celcius = %.1f fahrenheit", celcius, CelciusToFahrenheit(suhu(celcius)))
	fmt.Printf("\n%.0f celcius = %.2f kelvin", celcius, CelciusToKelvin(suhu(celcius)))

}