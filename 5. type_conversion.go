package main

import "fmt"

// Contoh konversi tipe data
func main() {
	var x int = 42
	var y float64 = float64(x) // Konversi int ke float
	var z int = int(y)         // Konversi float ke int

	fmt.Println("Integer to Float:", y)
	fmt.Println("Float to Integer:", z)
}
