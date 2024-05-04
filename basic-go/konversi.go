// Konversi tipe data
//
// - di go kadang kita butuh melakukan konversi tipe data dari satu tipe ke tipe lain
// - misal ingin konversi data dari int32 ke int63, dan lain2

// kode program konversi tipe data (1)

package main

import "fmt"

func main() {
	// nilai32 int32:= 32768
	// nilai64 int64 := int64(nilai32)
	// nilai16 int16 := int16(nilai32)

	var (
		nilai32 int32 = 32768
		nilai64 int64 = int64(nilai32)
		nilai16 int16 = int16(nilai32)
	)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
}
