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
	address1 := Address{"Padang", "Jagakarsa", "Jati"}
	address2 := address1 // copy value

	address2.City = "Jakarta Selatan"
	fmt.Println(address1) // tidak berubah
	fmt.Println(address2) // berubah menjadi Jakarta Selatan
}
