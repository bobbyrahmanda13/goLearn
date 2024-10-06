package main

import "fmt"

// step 1
// func main(){
//   var name string // jika type data pada name dihilangkan maka akan error
//
//   name = "Rahman Gaming"
//   fmt.Println(name)
//
//   name = "Rahman Ganteng"
//   fmt.Println(name)
// }

// step 2 = cara lansung karena golang tau kita mau menggunakan tipe data apa akan di detect secara otomatis
// func main() {

// var name string = "Rahman Gaming" // cara ini menjadi tidak wajib karena golang sudah tau type data yang digunakan jadi percuma bikin kek ginian
// 	var name = "Rahman Gaming"
// 	fmt.Println(name)
//
// 	name = "Rahman Ganteng"
// 	fmt.Println(name)
// }

// step 3 => pada golang kata kunci var tidak harus dibuat / tidak lah wajib, jadi bisa digantikan dengan ":=" saat menginisialisasikan data pada variable tersebut

// func main(){
// name := "Rahman Visalux" => ini cuman deklarasikan datanya saja jadi jika sudah dibuat maka tidak perlu membuat nya lagi jadi dibuat hanya untuk deklarasikan data jadi itu dibuat di awal saja
// name := "Rahman Visalux"
// fmt.Println(name)

// name = "Rahman Gaming" => jadi tidak diperbolahkan untuk dibuat lagi, jika misal dibuat lagi maka itu akan seperti variabel tersebut dibuat ulang / di deklarasikan ulang type  / variabel datanya
//   name = "Rahman Gaming"
//   fmt.Println(name)
// }

// Deklarasi Multiple Variable
// di golang kita bisa membuat banyak variabel sekaligus
// code yang dibuat akan lebih bagus dan mudah dibaca

func main() {
	name := "Rahman Visalux"
	fmt.Println(name)

	name = "Rahman Gaming"
	fmt.Println(name)

	var (
		firstName = "Rahman"
		lastName  = "Visalux"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
