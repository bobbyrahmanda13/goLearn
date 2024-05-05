// operasi matematika
// operator = keterangan
// + => penjumlahan
// - => pengurangan
// * => perkalian
// / => pembagian
// % => sisa pembagian

// note! => operasi matematika diatas digunakan pada type data number

package main

import "fmt"

func main() {
	var (
		// penjumlahan
		a = 10
		b = 10
		c = a + b
		// pengurangan
		d = 5
		e = 10
		f = d - e
		// perkalian
		kali1 = 50
		kali2 = 30
		tkali = kali1 * kali2

		// pembagian
		bagi1 = 20
		bagi2 = 5
		tbagi = bagi1 / bagi2

    // sisa pembagian
    sbagi1=50
    sbagi2=24
    tsbagi=sbagi1%sbagi2
	)

	fmt.Println(c)
	fmt.Println(f)
	fmt.Println(tkali)
	fmt.Println(tbagi)
	fmt.Println(tsbagi)
}
