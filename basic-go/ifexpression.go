// if Expression
// if adalah salah satu kata kunci yang digunakan untuk percabangan
// Percabangan artinya kita bisa mengeksekusi kode program tertentu ketika suatu kondisi terpenuhi
// Hampir di semua bahasa pemrograman mendukung if Expression

package main

import "fmt"

func main() {

  // var name string = "eko"
  name := "Jancok" // => type data string
  jancok := true // => type data boolean
  // jancok1 := 223 // => type data number

  if name == "Eko" {
		fmt.Println("Hello ", name)
	} else {
    fmt.Println("Hai boleh kenalan ")
  }

  if jancok == false {
    fmt.Println("Hy ", jancok)
  } else {
    fmt.Println("ini true")
  }

// Else Expression
// Blok if akan dieksekusi ketika kondisi if bernilai true 
// kadang kita ingin melakukan eksekusi program tertentu jika kondisi if bernilai false
// hal ini bisa dilakukan menggunakan else expression

}
