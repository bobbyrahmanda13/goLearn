/*
# Panic
- adalah function yang bisa kita gunakan untuk menghentikan program
- biasanya dipanggil ketika terjadi panic pada saat program kita berjalan
- saat panic function dipanggil, program akan terhenti, namun defer function
*/

package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	if !error {
		panic("ERROR")
	}
	/* if !error {
			panic("ERROR")
	} */
}

func main() {
	runApp(true)
}
