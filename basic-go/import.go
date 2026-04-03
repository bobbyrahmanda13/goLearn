package main

import (
	"basic-go/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("bobby")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // tidak bisa di akses (acces modifier)
	// fmt.Println(helper.sayGoodBye) // tidak bisa di akses (acces modifier)
}
