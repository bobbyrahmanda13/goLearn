// Type Declarations
// - type Declarations adalah kemampuan membuat ulang type data baru dari tipe data yg sudah ada
// - type Declarations biasanya digunakan untuk membuat alias terhadap tipe data yg sudah ada dengan tujuan agar lebih mudah dimengerti

// kode program type declarations

package main

import "fmt"

func main() {
	type NoKTP string // tipe data baru yg namanya NoKTP bertipe string tapi itu sebenarnya adalah alias untuk tipe data string

	var ktpRahman NoKTP = "223123123"

	var contoh string = "33333333"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpRahman)
	// fmt.Println(NoKTP("444444444"))
	// fmt.Println(contohKtp)
	fmt.Println(NoKTP(contohKtp))

}
