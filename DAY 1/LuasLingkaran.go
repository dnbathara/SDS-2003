package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung luas lingkaran
func LuasLingkaran(JariJari float64) float64 {
	return math.Pi * JariJari * JariJari
}

func main() {
	var JariJari float64

	// Input jari-jari lingkaran dari pengguna
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&JariJari)

	// Memanggil fungsi hitungLuasLingkaran
	luas := LuasLingkaran(JariJari)

	// Menampilkan hasil
	fmt.Printf("Luas lingkarannya adalah %.2f\n", luas)
}
