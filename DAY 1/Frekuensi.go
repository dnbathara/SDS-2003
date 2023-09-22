package main

import "fmt"

// Fungsi untuk menghitung frekuensi
func Frekuensi(gelombang int, waktu float64) float64 {
	return float64(gelombang) / waktu
}

func main() {
	var gelombang int
	var waktu float64

	fmt.Print("Masukkan banyak gelombang: ")
	fmt.Scan(&gelombang)
	fmt.Print("Masukkan waktu gelombang (s): ")
	fmt.Scan(&waktu)

	frekuensi := Frekuensi(gelombang, waktu)
	periode := 1 / frekuensi

	// Menampilkan hasil
	fmt.Printf("Frekuensi gelombang adalah %.2f Hz\n", frekuensi)
	fmt.Printf("Periode gelombang adalah %.2f s\n", periode)

}
