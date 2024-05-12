// Array
//
// Tipe data Array
// - Array adalah tipe data yang berisi kumpulan data dengan type yang sama
// - Saat membuat Array, kita perlu menentukan jumlah data yg bisa ditampung oleh Array tersebut
// - Daya tampung Array tidak bisa bertambah setelah Array dibuat
//
// Index di Array
//
// index | Data
// 0 | Rahman
// 1 | Visalux
// 2 | Gaming

package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Rahman"
	names[1] = "Visalux"
	names[2] = "Gaming"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// Membuat Array lansung
	// di Golang kita juga bisa membuat Array secara Lansung saat deklarasi variabel

	// contoh :

	var values = [3]int{
		90, 80, 95,
	}

	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(values)

}
