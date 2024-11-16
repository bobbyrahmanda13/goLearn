//# Returning Multiple Values

// - Function tidak hanya dapat mengembalikan satu value, tapi juga bisa multiple value
// - Untuk memberitahu jika function mengembalikan multiple value, kita harus menulis semua tipe data return value nya di function

package main

import "fmt"

func getFullName() (string, string) {
	return "Rahman", "Visalux"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}
