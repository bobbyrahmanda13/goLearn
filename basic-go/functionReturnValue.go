// # Function Return Value
// - function bisa mengembalikan data
// - Untuk memberitahu bahwa function mengambalikan data, kita harus menuliskan tipe data kembalian dari function tersebut
// - Jika function tersebut kita deklarasikan dengan tipe data pengembalian, maka wajib di dalam function nya kita harus mengembalikan data (menggunakan "return")
// - Untuk mengambalikan data dari function, kita bisa menggunakan kata kunci return, diikuti dengan datanya

package main

import "fmt"

func getHai(nama string) string {

  // return "Hai " + nama

  hello:="Hello " + nama

  return hello
}

func main(){
  // buat variabel result
  result := getHai("Rahman")
  fmt.Println(result)

  fmt.Println(getHai("Budi"))
  fmt.Println(getHai("joko"))
  fmt.Println(getHai("jancok"))

}


