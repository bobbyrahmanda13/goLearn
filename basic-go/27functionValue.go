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
func main() {
	goodbye := getGoodBye
	fmt.Println(goodbye("Rahman"))
}
