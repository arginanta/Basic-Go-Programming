package main

import "fmt"

// Contoh penggunaan variabel dan konstanta
func main() {
	var name string = "Argi"   // Variabel
	const pi float64 = 3.14159 // Konstanta

	fmt.Println("Name:", name)
	fmt.Println("Pi:", pi)

	name = "Sambadacamp" // Variabel bisa diubah
	fmt.Println("Updated Name:", name)
}
