package main

import "fmt"

func random() any {
	return true
}

func main() {
	var result any = random()

	// TEST:
	// var resultString string = result.(string)
	// fmt.Println(resultString)
	// var resultInt int = result.(int)
	// fmt.Println(resultInt)

	// NOTE: Type Assertions Menggunakan Switch
	// - saat salah menggunakan type assertions, maka bisa berakibat terjadi panic di aplikasi kita
	// - jika panic dan tidak ter-recover, maka otomatis program kita akan mati
	// - Agat lebih aman, sebaiknya kita menggunakan switch expression untuk melakukan type assertions

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)

	// TEST:
	// case bool:
	// 	fmt.Println("Boolean", value)

	default:
		fmt.Println("Unknown", value)
	}
}
