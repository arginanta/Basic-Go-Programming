package main

import "fmt"

// Contoh penggunaan map
func main() {
	// Map sederhana
	user := map[string]string{
		"name":  "Argi",
		"email": "argi@example.com",
	}
	fmt.Println("User:", user)

	// Map dalam map
	users := map[string]map[string]string{
		"user1": {"name": "Argi", "email": "argi@example.com"},
		"user2": {"name": "Siti", "email": "siti@example.com"},
	}
	fmt.Println("All Users:", users)
}
