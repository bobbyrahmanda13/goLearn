/*
# Operasi perbandingan

- operasi perbandingan adalah operasi untuk membandingkan dua buah data
- operasi perbandingan adalah operasi yg menghasilkan nilai boolean ( benar / salah )
- jika hasl operasinya adalah benar, maka nilainya adalah true
- jika hasl operasinya adalah salah, maka nilainya adalah false
*/

/*
# Tabel Operator Perbandingan

Operator | Keterangan
> | Lebih dari
< | Kurang dari
>= | Lebih dari Sama Dengan
<= | Kurang dari Sama Dengan>
== | Sama Dengan
!= | Tidak Sama Dengan
*/

/*
note:
- operator > dan < bisa dibandingkan / digunakan untuk angka tapi tidak bisa digunakan untuk type data number ataupun boolean
- namun operator  ( == dan != ) hanya bisa digunakan untuk type data string, jadi menggunakannya bisa untuk membandingan type data string ( misal untuk membandingan string itu sama atau tidak )
- namun untuk tipe data number kita bisa menggunakan semua operator perbandingan
- tapi bisa menggunakan operator yang < dan > tapi itu berdasarkan huruf abjad
jadi kalau misal B > A = True jika B < A = false tapi jarang sekali yang menggunakan ini
*/

// contoh :

package main

import "fmt"

func main() {
	var (
		name1 = "rahman"
		name2 = "rahman"

		name3 = "jancok"
		name4 = "taik"

		name5 = "jancok"
		name6 = 6
		name7 = 10

		result  bool = name1 == name2
		result2 bool = name3 == name4
		result3 bool = name3 != name4
		result4 bool = name3 != name5
		result5 bool = name6 >= name7
		result6 bool = name6 <= name7
	)

	fmt.Println(result)  // true
	fmt.Println(result2) // false
	fmt.Println(result3) // true
	fmt.Println(result4) // false
	fmt.Println(result5) // false
	fmt.Println(result6) // true

	var (
		nilaiAkhir = 0
		absensi  = 80
	)
	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus)

}
