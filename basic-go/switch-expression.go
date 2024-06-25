// Switch Expression
// - Selain if expression, untuk melakukan percabangan, kita juga bisa menggunakan Switch Expression
// - Switch expression sangat sederhana dibandingan if
// - Biasanya switch expression digunakan untuk melakukan pengecekan ke kondisi dalam satu variable

package main

import "fmt"

func main() {

	name := "Rahman"

	switch name {

	case "Rahman":
		// fmt.Println("Hello ", name)
		fmt.Println("Hello Rahman")

	case "Richo":
		// fmt.Println("Hello ", name)
		fmt.Println("Hello Richo")

	default: // ini sama dengan else pada if expression
		fmt.Println("Hi, Boleh Kenalan ?")

	}

}
