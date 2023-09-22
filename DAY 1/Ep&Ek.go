package main

import "fmt"

// Fungsi untuk menghitung luas segitiga
func EnergiPotensial(M float64, V float64) float64 {
	return 0.5 * M * V * V
}

func EnergiKinetik(Massa float64, Gravitasi float64, Heigh float64) float64 {
	return Massa * Gravitasi * Heigh
}

func main() {
	var M, V, Massa, Gravitasi, Heigh float64

	fmt.Println("Rumus Menghitung Energi Kinetik")
	fmt.Print("Masukkan besar massanya: ")
	fmt.Scan(&M)
	fmt.Print("Masukkan besar kecepatan-nya: ")
	fmt.Scan(&V)
	// Memanggil fungsi Energi Kinetik
	EP := EnergiPotensial(M, V)
	fmt.Printf("Energi Potensialnya adalah %.2f\n\n", EP)

	fmt.Println("=====================================\n\n")
	fmt.Println("Rumus Menghitung Energi Potensial")

	fmt.Println("Masukkan besar massanya: ")
	fmt.Scan(&Massa)
	fmt.Println("Masukkan besar ravitasinya: ")
	fmt.Scan(&Gravitasi)
	fmt.Println("Masukkan besar ketinggiannya: ")
	fmt.Scan(&Heigh)
	//Memanggil fungsi Energi Potensial
	EK := EnergiKinetik(Massa, Gravitasi, Heigh)
	fmt.Println("Energi Kinetiknya adalah :", EK)

}
