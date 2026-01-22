// # Break & Continue
// - Break dan Continue adalah kata kunci yg bisa digunakan dalam perulangan
// - Break digunakan untuk menghentikan seluruh perulangan
// - Continue adalah digunakan untuk menghentikan perulangan yang berjalan, dan lansung melanjutkan ke perulangan selanjutnya

package main

import "fmt"

func main() {

	// fmt.Println("hello world")
	// # Kode Program Break
	for i := 0; i < 10; i++ {
		if i == 7 {
			break
		}
		fmt.Println("Perulangan Ke", i)

	}

	fmt.Println(" ")

	// # Kode Program Continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan Ke", i)
	}

}
