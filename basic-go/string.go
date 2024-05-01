package main

import "fmt"

// func main(){
//   fmt.Println("Rahman")
//   fmt.Println("Rahman RVBR")
//   fmt.Println("Rahman RVBR Gaming")
// }

// function untuk string
// len("string") = menghitung jumlah karakter di string
// "string"[number] = mengambil karakter pada posisi yang ditentukan

func main(){
  fmt.Println(len("Rahman"))
  fmt.Println("Rahman RVBR"[0]) /* =>  hasil yang di dapat dalam bentuk byte */
  fmt.Println("Rahman RVBR Gaming"[1]) /* =>  hasil yang di dapat dalam bentuk byte */
}
