package main

import "fmt"

// Contoh for loop dan penggunaan break/continue
func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // Lewati iterasi ke-5
		}
		if i == 8 {
			break // Keluar dari loop saat i = 8
		}
		fmt.Println("Value:", i)
	}
}
