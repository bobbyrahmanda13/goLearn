package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Bobby Rahmanda", "Bobby"))
	fmt.Println(strings.Split("Bobby Rahmanda", " "))
	fmt.Println(strings.ToLower("Bobby Rahmanda"))
	fmt.Println(strings.ToUpper("Bobby Rahmanda"))
	fmt.Println(strings.ReplaceAll("Bobby Rahmanda Bobby Jelek", "Bobby", "Riko"))
}
