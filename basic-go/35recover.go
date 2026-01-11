/*
# Recover
- Recover adalah function yang bisa kita gunakan untuk menangkap data panic
- dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan
*/

package main

import "fmt"

func endAppp() {
	fmt.Println("End App")

	// # Recover yang benar

	message := recover()
	fmt.Println("Terjadi Error", message)
}

// # Recover yang salah

func runAppp(error bool) {
	defer endAppp()
	if error {
		panic("ERROR")
	}

	// message := recover()
	// fmt.Println("Terjadi Error", message)
}

func main() {
	runAppp(true)
}
