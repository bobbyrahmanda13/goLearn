/* Multiple Main Function

- di Golang, function dalam module / projek adalah unik, artinya tidak boleh membuat nama function yang sama
- Oleh karena itu, jika kita membuat file baru, misal sample.go, lalu membuat nama function yang sama yaitu main
- Maka kita tidak bisa melakukan build module, karena main function tersebut duplikat dengan yang ada di main function helloword.go */

package main

import "fmt"

func main() {
	fmt.Println("sample go")
}
