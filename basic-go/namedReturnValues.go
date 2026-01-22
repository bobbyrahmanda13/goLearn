// # Named Return Values

// - Biasanya saat kita memberi tahu bahwa sebuah function mengembalikan value, maka kita hanya mendeklarasikan tipe data return value di function
// - Namun kita juga bisa membuat variable secara lansung di tipe data function nya

package main

import "fmt"

func getCompleteName() (firstName string, middleName string, lastName string) {
	firstName = "Richo"
	middleName = "Hardiansyah"
	lastName = "Buncit"

	return firstName, middleName, lastName
}

func main() {

	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
