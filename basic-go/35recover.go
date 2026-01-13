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

	// testing
	message := recover()
	fmt.Println("Terjadi Error", message)
}

func asalApp() {
	fmt.Println("end appp coy")

	messagess := recover()
	fmt.Println("Terjadi Error cuy", messagess)
}

func jalanApp(taik string) {
	defer asalApp()
	if taik == "anjir" {
		panic("Error boss")
	}
}

func runAppp(error bool) {
	defer endAppp()
	if error {
		panic("ERROR")
	}

	// # Recover yang salah

	// message := recover()
	// fmt.Println("Terjadi Error", message)
}

func main() {
	runAppp(true)
	jalanApp("anjir")
	fmt.Println("lanjut programmnya")
}
