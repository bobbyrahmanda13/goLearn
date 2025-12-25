// Operasi Boolean
// Operator | Keterangan
// - && | Dan
// - || | Atau
// - !  | Kebalikan

// Operasi && = dan / AND

// Nilai 1 | Operator | Nilai 2 | Hasil
// -------------------------------------
// true    | &&       | true    | true
// true    | &&       | false   | false
// false   | &&       | true    | false
// false   | &&       | false   | false
// -------------------------------------
// Note !! => jika salah satunya itu false maka hasil nya adalah false

// Operasi || = atau / OR

// Nilai 1 | Operator | Nilai 2 | Hasil
// -------------------------------------
// true    | ||       | true    | true
// true    | ||       | false   | true
// false   | ||       | true    | true
// false   | ||       | false   | false
// -------------------------------------
// Note !! => jika salah satu nya itu true maka hasil nya adalah true

// Operasi ! = kebalikan

// Operator | Nilai 2 | Hasil
// -------------------------------------
// !       | true    | true
// !       | false   | true
// -------------------------------------

// Kode Program Operasi Boolean

package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80 // true
	var lulusAbsensi bool = absensi > 80       // false

	var lulus bool = lulusNilaiAkhir && lulusAbsensi // false

	var lulus2 bool = lulusNilaiAkhir || lulusAbsensi // true

	fmt.Println(lulus)  // false
	fmt.Println(lulus2) // true
}
