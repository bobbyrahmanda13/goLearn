/*
# Function Value

- Function adalah first class citizen
- Function juga merupakan tipe data, dan bisa disimpan di dalam variable

*/

package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name

}

func sayangSaranghaee(name string) string {
	return "Sayang " + name
}

func main() {
	goodbye := getGoodBye
	fmt.Println(goodbye("Rahman"))

	contoh := sayangSaranghaee
	misal := sayangSaranghaee

	fmt.Println(contoh("Rahman"))
	fmt.Println(misal("Rahman"))

}
