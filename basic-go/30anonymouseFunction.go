/*
# Anonymouse Function
- Sebelumnya setiap membuat function, kita akan selalu memberikan sebuah nama function tersebut
- Namun kadang ada kalanya lebih mudah membuat function secara lansung di variable atau parameter tanpa harus membuat function terlebih dahulu
- Hal tersebut dinamakan Anonymous Function atau function tanpa nama
*/
package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}
	registerUser("bobby", blacklist)

	// registerUser("anjing", func(name string) bool {
	// 	return name == "anjing"
	// })
}
