package main

import "fmt"

// Fungsi untuk menghitung luas segitiga
func LuasSegitiga(alas float64, tinggi float64) float64 {
	return 0.5 * alas * tinggi
}

func main() {
	var alas, tinggi float64

	fmt.Print("Masukkan panjang alas: ")
	fmt.Scan(&alas)
	fmt.Print("Masukkan tingginya: ")
	fmt.Scan(&tinggi)

	// Memanggil fungsi hitungLuasSegitiga
	luas := LuasSegitiga(alas, tinggi)

	fmt.Printf("Luas segitiga-nya adalah %.2f\n", luas)
}
