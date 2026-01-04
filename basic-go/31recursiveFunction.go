/*
# Recursive Function
- Recursive Function adalah function yang memanggil function dirinya sendiri
- Kadang dalam pekerjaan, kita sering menemui kasus dimana menggunakan recursive function lebih mudah dibandingkan tidak menggunakan recursive function
- Contoh Kasus yang lebih diselesaikan menggunakan recursive adalah Factorial
*/

package main

import "fmt"

// - Menggunakan Factorial For Loop

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func main() {
	fmt.Println(factorialLoop(10))
}
