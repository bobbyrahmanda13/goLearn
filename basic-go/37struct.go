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

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
}
