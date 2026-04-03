package main

import (
	"fmt"

	"basic-go/helper"
)

func main() {
	result := helper.SayHello("bobby")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // tidak bisa di akses (acces modifier)
	// fmt.Println(hekkkkkkkkajlper.sayGoodBye) // tidak bisa di akses (acces modifier)
}
