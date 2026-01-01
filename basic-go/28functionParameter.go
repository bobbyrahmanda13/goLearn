/*
# Function as Parameter (sebagai parameter)
- function tidak hanya bisa kita simpan di dalam variable sebagai value
- namun juga bisa kita gunakan sebagai parameter untuk function lain
*/

package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Babi" {
		return "..."
	}
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
	sayHelloWithFilter("Rahman", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Tai", filter)
}
