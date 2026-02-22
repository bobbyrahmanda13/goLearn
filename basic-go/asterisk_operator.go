// NOTE:
// * Operator *
// - Saat kita mengubah variable pointer, maka yang barubah hanya variable tersebut,
// - semua varible yang mengacu ke data yang sama tidak akan berubah
// - jika kita ingin mengubah seluruh variable yang mengacu ke data tersbut, kita bisa menggunakan operator *

package main

import "fmt"

type Addres struct {
	City, Province, Country string
}

func main() {
	var address1 Addres = Addres{"Padang", "Jagakarsa", "Jati"}
	var address2 *Addres = &address1 // dengan pointer

	// TEST:
	// address1 := Addres{"Padang", "Jagakarsa", "Jati"}
	// address2 := &address1 // dengan pointer

	address2.City = "Jakarta Selatan"

	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah menjadi Jakarta Selatan

	address2 = &Addres{"Bukittinggi", "Jakarta Selatan", "Indonesia"}
}
