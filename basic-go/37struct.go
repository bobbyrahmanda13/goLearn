/*
# Struct
- Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
- Struct biasanya representasi data dalam program aplikasi yang kita buat
- Data di Struct disimpan dalam field
- Sederhananya struct adalah kumpulan dari field
*/

/*
Membuat data Struct
- Struct adalah template data atau prototype data
- Struct tidak bisa langsung digunakan
- Namun Kita bisa membuat data / object dari struct yang telah kita buat
*/

package main

import "fmt"

type Customer struct {
	// boleh seperti ini
	Name, Address string
	Age           int

	/*
	   // boleh juga seperti ini
	   Name    string
	   Address string
	   Age     int
	*/
}

func main() {
	var bobby Customer

	fmt.Println(bobby)

	bobby.Name = "Bobby Rahmanda"
	bobby.Address = "Indonesia"
	bobby.Age = 30
	fmt.Println(bobby)
	fmt.Println(bobby.Name)
	fmt.Println(bobby.Address)
	fmt.Println(bobby.Age)

	/*
	   # Struct Literals
	   - sebelumnya kita telah membuat data dari struct, namun sebenernya ada banyak cara yang bisa kita gunakan untuk membuat data dari struct
	*/

	bulek := Customer{
		Name:    "Bulek",
		Address: "Indonesia",
		Age:     30,
	}
	fmt.Println(bulek)
	riko := Customer{"Budi", "Indonesia", 30}
	fmt.Println(riko)

	/*
		# Struct Method
		- struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function
		- Namun jika kita ingin menambahkan method ke dalam struct, sehingga sekan-akan sebuah struct memiliki function
		- Method adalah function
	*/
}
