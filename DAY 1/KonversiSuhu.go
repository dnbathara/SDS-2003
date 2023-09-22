package main

import "fmt"

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}
func celsiusToReamur(celsius float64) float64 {
	return celsius * 4 / 5
}
func celciusToKelvin(celcius float64) float64 {
	return celcius + 273.15
}

func main() {
	var celsius float64

	// Input suhu dalam Celsius dari pengguna
	fmt.Print("Masukkan suhu dalam Celsius: ")
	fmt.Scan(&celsius)

	// Menggunakan fungsi celsiusKeFahrenheit untuk konversi
	fahrenheit := celsiusToFahrenheit(celsius)

	// Menggunakan fungsi celsiusKeReamur untuk konversi
	reamur := celsiusToReamur(celsius)

	kelvin := celciusToKelvin(celsius)

	// Menampilkan hasil konversi
	fmt.Printf("%.2f Celsius sama dengan %.2f Fahrenheit\n", celsius, fahrenheit)
	fmt.Printf("%.2f Celsius sama dengan %.2f Reamur\n", celsius, reamur)
	fmt.Printf("%.2f Celsius sama dengan %.2f Kelvin\n", celsius, kelvin)
}
