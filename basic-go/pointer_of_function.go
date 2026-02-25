// NOTE:
// * Pointer di function
// - Saat kita membuat parameter di function, secara default adalah pass by value, artinya data akan di copy lalu dikirim ke function tersebut
// - Oleh karena itu, jika kita mengubah data di dalam function, data yang aslinya tidak akan pernah berubah
// - Hal ini membuat variable menjadi aman, karena tidak akan bisa diubah
// - Namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut
// - Untuk melakukan ini, kita juga bisa menggunakan pointer di function
// - Untuk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator * di parameternya

package main

import "fmt"

type Alamatt struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(alamat *Alamatt) {
	alamat.Country = "Indonesia"

	// TEST:
	// fmt.Println(alamat)
}

func main() {
	var alamat *Alamatt = &Alamatt{"Medan", "Sumatera Utara", ""}
	ChangeCountryToIndonesia(alamat)

	fmt.Println(alamat) // tidak berubah
}
