package main

import "fmt"

// Contoh array dan slice
func main() {
	// Array
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr)

	// Slice
	slice := arr[1:4]
	fmt.Println("Slice:", slice)

	// Tambahkan elemen ke slice
	newSlice := append(slice, 6)
	fmt.Println("New Slice:", newSlice)
}
