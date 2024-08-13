// For Expression (For Loops)

// For :
// Dalam bahasa pemrograman, biasanya ada fitur yang bernama perulangan
// salah satu fitur perulangan adalah for loops

// Kode Program For
package main

import (
	"fmt"
)

func main() {
	// buat variable baru
	counter := 1

  // lakukan perulangan dengan for dimana counter jika kurang atau sama dengan 10 maka cetak print Perulangan ke , counter +1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++

	}

}
