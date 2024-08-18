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
	fmt.Println("Selesai")

	//# For dengan Statement
	//  Dalam for, kita bisa menambahkan statement, dimana terdapat 2 statement yang bisa ditambahkan di for
	// * init statement, yaitu statement sebelum for di eksekusi
	// * Post statement, yaitu statement yang akan selalu di eksekusi di akhir tiap perulangan

	//# kode program For dengan statement
	for counter2 := 1; counter2 <= 10; counter2++ {
		fmt.Println("Perulangan counter2 ke", counter2)
	}

	fmt.Println("Selesai")

	// # For Range
	// for bisa digunakan untuk melakukan iterasi terhadap semua data collection
	// data collection contoh nya Array, Slice dan Map

	// % Manual
	// names := []string{"Dassa", "Putra", "Ganteng"}
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// % For Range

	names := []string{"Dassa", "Putra", "Ganteng"}

	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	for i, name2 := range names {
		fmt.Println("index", i, "=", name2)
	}

	// # jika tidak butuh index nya maka diganti dengan _ (underscore)
	for _, name3 := range names {
		// fmt.Println("Urutan nama Perkata", name3)
		fmt.Println(name3)
	}

}
