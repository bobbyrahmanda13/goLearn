// Tipe Data Map
// - Pada Array or Slice, untuk mengakses data, kita menggunakan index Number dimulai dari 0
// - Map adalah tipe data lain yg berisikan kumpulan data yg sama, namun kita bisa menentukan jenis tipe data index yg akan kita gunakan
// - sederhananya, Map adalah tipe data kumpulan key-value (kata kunci - nilai), dimana kata kuncinya bersifat unik, tidak boleh sama
// - Berbeda dengan Array dan Slice, jumlah data yg kita masukkan ke dalam Map boleh sebanyak-banyaknya, asalkan kata kunci nya berbeda, jika kita gunakan kata kunci sama, maka secara otomatis data sebelumnya akan diganti dengan data baru

package main

import "fmt"

func main(){

  // # ini map yang lansung dimasukkan datanya
  person := map[string]string{ // map[string] = ini adalah map type value nya string, sedangkan string{} ini adalah type untuk nilai nya
    "name" : "Rahman",
    "address" : "Paris",
  }

  // 1. deklarasikan map
  // var person map[string]string = map[string]string{} // ini adalah map kosong alias tanpa data

  // 2. lalu isikan datanya person["name"] adalah lah key nya
  // sedangkan setelah = itu adalah value data nya
  // person["name"] = "Rahman"
  // person["address"] = "Paris"

  fmt.Println(person)
  fmt.Println(person["name"])
  fmt.Println(person["address"])


}
