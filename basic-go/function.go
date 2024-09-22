// # Function

// - sebelumnya kita sudah mengenal sebuah function yang wajib dibuat agar program kita bisa berjalan, yaitu function main()
// - Function adalah sebuah blok kode yang sengaja dibuat dalam program agar bisa digunakan berulang-ulang
// - Cara membuat function sangat sederhana, hanya dengan menggunakan kata kunci "func" lalu diikuti dengan nama function nya dan blok kode isi function nya
// - Setelah membuat function, kita bisa mengeksekusi function tersebut dengan memanggilnya menggunakan kata kunci nama function nya diikuti tanda kurung buka "{", dan kurung tutup "}"

package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func main() {
	sayHello()
  sayHelloTo("Rahman","Visalux", 500, true)

  // ! Note : sebuah function bisa di panggil sebanyak2nya atau bisa dibilang tidak terbatas
}

//* Function Parameter 
// - Saat membuat function, kadang-kadang kita membutuhkan data dari luar, atau kita sebut parameter.
// - Kita bisa menambahkan parameter di function, bisa lebih dari satu.
// - Parameter tidak lah wajib, jadi kita bisa membuat function tanpa parameter seperti sebelumnya yang sudah kita buat
// - Namun jika kita menambahkan parameter di function, maka ketika memanggil function tersebut, kita wajib memasukkan data ke parameternya

func sayHelloTo(firstName string, lastName string, nomor int, betul bool){
  fmt.Println("Hello", firstName, lastName, nomor, betul)
}

