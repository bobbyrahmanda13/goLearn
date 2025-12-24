// Switch Expression
// - Selain if expression, untuk melakukan percabangan, kita juga bisa menggunakan Switch Expression
// - Switch expression sangat sederhana dibandingan if
// - Biasanya switch expression digunakan untuk melakukan pengecekan ke kondisi dalam satu variable

package main

import "fmt"

func main() {

	name := "Rahman"

	switch name {

	case "Rahman":
		// fmt.Println("Hello ", name)
		fmt.Println("Hello Rahman")

	case "Richo":
		// fmt.Println("Hello ", name)
		fmt.Println("Hello Richo")

	default: // ini sama dengan else pada if expression
		fmt.Println("Hi, Boleh Kenalan ?")

	}

	// Switch dengan Short Statement
	// - Sama dengan If, Switch juga mendukung short statement sebelum variable yg akan di check kondisinya
	// example
	jancok := "Taik"

	switch length := len(jancok); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// note: untuk default sendiri bisa ditambahkan, bisa juga tidak

	// Swith Tanpa Kondisi
	// - Kondisi di switch expression tidak wajib
	// - jika kita menggunakan kondisi di switch expression, kita bisa menambahkan kondisi tersebut di setiap case nya

	// Kode program Switch Tanda Kondisi
	name = "taikkkkkkkkkk"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}

}
