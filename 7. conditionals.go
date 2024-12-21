package main

import "fmt"

// Contoh if dan switch
func main() {
	age := 20

	// If statement
	if age < 18 {
		fmt.Println("You are a minor.")
	} else if age >= 18 && age < 65 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a senior citizen.")
	}

	// Switch statement
	grade := "A"
	switch grade {
	case "A":
		fmt.Println("Excellent!")
	case "B":
		fmt.Println("Good job.")
	default:
		fmt.Println("Keep trying.")
	}
}
