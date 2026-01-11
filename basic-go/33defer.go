/* # Defer
- Defer function adalah function yang bisa kita jadwalkan untuk di eksekusi setelah sebuah function selesai di eksekusi
- Defer function akan selalu di eksekusi walaupun terjadi error di function yang di eksekusi */

package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging()
	fmt.Println("Run Application")
	/*
		// error
			logging()
	*/
}

func main() {
	runApplication()
}
