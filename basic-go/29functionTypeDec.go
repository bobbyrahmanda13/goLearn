/*
	Function Type Declaration

- Kadang jika function terlalu panjang,  agak ribet untuk menuliskannya di dalam parameter
- Type Declarations juga bisa digunakan membuat alias function, sehingga akan mempermudah kita menggunakan function sebagai parameter
*/
package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilters(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func spamFilters(name string) string {
	if name == "Anjing" {
		return "..."
	}
	if name == "Tai" {
		return "..."
	}
	if name == "Asu" {
		return "..."
	}
	return name
}

func main() {
}
