package main

import "fmt"

// Fungsi untuk menghitung luas segitiga
func LuasPersegi(sisi float64) float64 {
	return sisi * sisi
}

func main() {
	var sisi float64

	fmt.Print("Masukkan panjang sisinya: ")
	fmt.Scan(&sisi)

	// Memanggil fungsi hitungLuasSegitiga
	luas := LuasPersegi(sisi)

	fmt.Printf("Luas segitiga-nya adalah %.2f\n", luas)
}
