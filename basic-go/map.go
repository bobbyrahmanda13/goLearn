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

  // note! pada map ini bebas untuk buat key nya tanpa ada batasannya

  // 1. deklarasikan map
  // var person map[string]string = map[string]string{} // ini adalah map kosong alias tanpa data

  // 2. lalu isikan datanya person["name"] adalah lah key nya
  // sedangkan setelah = itu adalah value data nya
  // person["name"] = "Rahman"
  // person["address"] = "Paris"

  fmt.Println(person["name"])
  fmt.Println(person["address"])
  fmt.Println(person)

  fmt.Println(person["salah"])


  // * Function Map
  // Operasi | Keterangan
  // len(map) | Untuk mendapatkan jumlah data di map
  // map[key] | Mengambil data di map dengan key
  // map[key] = value | Mengubah data di map dengan key
  // make(map[TypeKey]TypeValue) | Membuat map baru
  // delete(map, key) | Menghapus data di map dengan key

}
