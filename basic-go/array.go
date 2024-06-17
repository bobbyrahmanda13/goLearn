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

	// Function Array
	// Operasi | Keterangan
	// len(array) | Untuk mendapatkan panjang Array
	// array[index] | Mendapatkan data di posisi index
	// array[index] = value | mengubah data di posisi index

	// contoh kode Program array
	var valuess = [...]int{ // [...] tanda ini adalah untuk menampung berapa banyak array yang ingin dimasukkan, jadi jika kita tidak tau ingin memasukkan jumlah data array nya berapa maka buat saja tanda itu
		//! note menggunakan tanda [...] datanya harus di deklarasikan dahulu jika tidak akan error
		90,
		80,
		95,
		100,
		120,
	}

	fmt.Println(valuess)
	fmt.Println("total banyak data dalam array :", len(valuess))
	valuess[0] = 100 // format valuess[0] = 100 => valuess adalah array, [0] ini = index, 100 adalah value
	fmt.Println(valuess)

	// note : question => bagaimana cara menghapus data array
	// answer : digolang tidak ada cara untuk menghapus data array
	// jadi jika array nya sudah ditentukan ya akan seperti yang ditentukan itu tidak dapat di hapus, palingan index nya di isi menjadi kosong tapi array nya tetap
}
