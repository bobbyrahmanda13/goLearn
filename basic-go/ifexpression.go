// if Expression
// if adalah salah satu kata kunci yang digunakan untuk percabangan
// Percabangan artinya kita bisa mengeksekusi kode program tertentu ketika suatu kondisi terpenuhi
// Hampir di semua bahasa pemrograman mendukung if Expression

package main

import "fmt"

func main() {

	// var name string = "eko"
	name := "Jancok" // => type data string
	jancok := true   // => type data boolean
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

	// Else If Expression
	// - Kadang dalam if, kita butuh membuat beberapa kondisi
	// - Kasus seperti ini, kita bisa menggunakan else if expression
	// - jika kondisi if dan else if tidak terpenuhi maka dapat menambahkan else

	// # Kode program if Else

	nama := "Riko"

	if nama == "Rahman" {
		fmt.Println("Hello", nama)
	} else if nama == "Riko" {
		fmt.Println("Hello", nama)
	} else {
		fmt.Println("Hai Boleh Kenalan")
	}

  // IF dengan short statement
  // if mendukung short statement sebelum kondisi 
  // hal ini sangat cocok untuk membuat statement yang sederhana sebelum melakukan pengecekan terhadap kondisi

  // Kode Program if short statement

  // biasanya seperti ini :

  // taik := "visalux" 
  // masha := len(taik)
  //
  // if masha > 5 {
  //   fmt.Println("Nama Terlalu panjang")
  // } else {
  //   fmt.Println("Nama sudah benar")
  // }

  // Kode Program if short statement

  taik := "visalux" 

  if masha := len(taik); masha > 5 {
    fmt.Println("Nama Terlalu panjang")
  } else {
    fmt.Println("Nama sudah benar")
  }

  // cat :
  // masha adalah variabel, len(taik) = statement len menghitung jumlah karakter, taik = variabel yang berisi string  kata "visalux", ; (titik koma) = sebagai pembatas untuk kondisi
  // masha > 5 => maksudnya variabel masha yang berisi panjang karakter dari variabel taik, lalu > 5 yang jika angka nya lebih dari 5 maka tercetak "Nama Terlalu Panjang", jika tidak terpenuhi (else) maka yang tercetak adalah "Nama sudah benar"



}

