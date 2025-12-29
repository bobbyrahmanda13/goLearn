/* Variadic Function

- parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varargs(variabel argumen)
- Varags artinya datanya bisa menerima lebih dari satu input, atau di anggap saja semacam array
- Bedanya dengan parameter biasa dengan type data Array :
  - Jika parameter tipe Array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function
  - Jika parameter menggunakan varags, kita bisa lansung mengirim datanya, jika lebih dari satu, cukup gunakan tanda koma
*/

package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {

	fmt.Println(sumAll(10, 10, 10, 10, 10))

	fmt.Println(sumAll(10, 10, 10, 10, 10, 10, 10, 10))
	fmt.Println(sumAll(10, 10, 10))

	/*
		Slice parameter
		- kadang ada kasus dimana kita sudah menggunakan variadic function namun memiliki berupa slice
		- kita bisa menjadikan slice sebagai varags paremeter
	*/

	numer := []int{10, 10, 10, 10, 10}
	fmt.Println(sumAll(numer...))

}
