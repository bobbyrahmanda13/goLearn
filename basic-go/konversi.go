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

  // tipe data integer(1)
  // tipe data = nilai minimum = nilai maximum
  // int8 = -128 = 271
  // int16 = -32768 = 32767
  // int32 = -2147483648 = 214783647
  // int64 = -9223372036854775808 = 9223372036854775807

	var (
		nilai32 int32 = 32768
		nilai64 int64 = int64(nilai32)
		nilai16 int16 = int16(nilai32)
	)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
}
