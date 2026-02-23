package main

import "fmt"

// NOTE:
// * Operator New
// - Sebelumnya untuk membuat pointer dengan menggunakan Operator "&"
// - Golang juga memiliki function new yang bisa dignakan untuk membuat pointer
// - Namun function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal

type Addresss struct {
	City, Province, Country string
}

func main() {
	var alamat1 *Addresss = &Addresss{} // *Address adalah pointer ke struct
	var alamat2 *Addresss = alamat1

	alamat2.Country = "Indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
