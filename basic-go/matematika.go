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
		sbagi1 = 50
		sbagi2 = 24
		tsbagi = sbagi1 % sbagi2
	)

	a += 10 // a = a + 10
	d -= 10 // d = d - 10

	kali1 *= 10  // kali1 = kali1 * 10
	bagi1 /= 10  // bagi1 = bagi1 / 10
	sbagi1 %= 10 // sbagi1 = sbagi1 % 10

	fmt.Println(a)
	fmt.Println(d)
	fmt.Println(kali1)
	fmt.Println(bagi1)
	fmt.Println(sbagi1)
	fmt.Println("--------------------")

	fmt.Println(c)
	fmt.Println(f)
	fmt.Println(tkali)
	fmt.Println(tbagi)
	fmt.Println(tsbagi)

	// augmented assigment
	// operasi matematika => augmented assigment

	// a = a + 10 => a += 10
	// a = a - 10 => a -= 10
	// a = a * 10 => a *= 10
	// a = a / 10 => a /= 10
	// a = a % 10 => a %= 10

	// Unary Operator
	// operator = keterangan
	//   ++ => a = a + 1
	//   -- => a = a - 1
	//   - => Negative
	//   + => Positive
	//   ! => Boolean Kebalikan

	// Kode Program Unary Operator
	fmt.Println("-----------------------")

	var j = 1
	fmt.Println(j)

	j++ // j = j +1
	fmt.Println(j)

	j++ // j = j +1
	fmt.Println(j)
  fmt.Println("----------------------")

  j--
  fmt.Println(j)

}
