// constant
package main

/*
- adalah variable yang nilai nya tidak bisa diubah lagi setelah pertama kali diberi nilai
- cara pembuatannya constant mirip dengan variable, yg membedakannya hanya kata kunci yang digunakan adalah const, bukan var
- saat pembuatan constant, kita wajib lansung menginisialisasikan datanya
*/
func main() {
	/*
		cara biasa
		const firstname string = "rahman"
		const lastname = "gaming"
	*/

	// Deklarasi ultiple constant => membuat constant dengan sekaligus banyak

	const (
		firstName string = "Rahman"
		lastName         = "juancok"
	)
	/*
	   error
	   firstName = "Tidak bisa diubah"
	   lastName = "Tidak bisa diubah"
	*/
}
