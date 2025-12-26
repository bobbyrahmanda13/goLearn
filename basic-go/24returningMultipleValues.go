//# Returning Multiple Values

// - Function tidak hanya dapat mengembalikan satu value, tapi juga bisa multiple value
// - Untuk memberitahu jika function mengembalikan multiple value, kita harus menulis semua tipe data return value nya di function

package main

import "fmt"

func getFullName() (string, string) {
	return "Rahman", "Visalux"
}

func getGameName() (string, string) {
	return "GOD", "Visalux"
}

func getNameSayang() (string, int, string, bool) {

	return "Windy Annisa", 17, "Juni", true
}

func main() {

	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName) // butuh semuanya

	// # Menghiraukan Return Value
	// - Multiple return value wajib ditangkap value nya
	// - Jika kita menghiraukan return value tersebut, kita bisa menggunakan tanda _ (garis bawah)

	_, namaAkhir := getGameName()
	fmt.Println(namaAkhir)

	namaAwal, _ := getGameName()
	fmt.Println(namaAwal)

	NamaPasangan, tglLahir, BulanLahir, benar := getNameSayang()
	fmt.Println(NamaPasangan, tglLahir, BulanLahir, benar)

	NamaCewek, tgl, bulan, _ := getNameSayang()
	fmt.Println(NamaCewek, tgl, bulan)
}
