package main

import "fmt"

// NOTE:
// * Pass by Value
// - Secara default di Go-Lang semua variable itu di passing by value, bukan by Reference
// - Artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain, sebenarnya yang dikirim adalah value nya

type Address struct {
	City, Province, Country string
}

func main() {
	// TEST:
	// address1 := Address{"Padang", "Jagakarsa", "Jati"}
	// address2 := address1 // copy value // default
	// address2.City = "Jakarta Selatan"
	// fmt.Println(address1) // tidak berubah = {Padang Jagakarsa Jati}
	// fmt.Println(address2) // berubah menjadi Jakarta Selatan = {Jakarta Selatan Jagakarsa Jati}

	// NOTE:
	// * Pointer
	// - adalah kemampuan membuat Reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada
	// - sederhananya, dengan kemampuan pointer, kita bisa membuat pass by reference (data yang sama)

	// NOTE:
	// * Operator &
	// - untuk membuat sebuah variable dengan nilai pointer ke variable yang lain, kita bisa menggunakan operator & di ikuti dengan nama variable nya

	var address1 Address = Address{"Padang", "Jagakarsa", "Jati"}
	var address2 *Address = &address1 // dengan pointer

	// address1 := Address{"Padang", "Jagakarsa", "Jati"}
	// address2 := &address1 // dengan pointer

	address2.City = "Jakarta Selatan"
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah menjadi Jakarta Selatan
}
