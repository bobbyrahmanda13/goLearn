// # Type data Slice
// type data ini adalah tipe data yang sangat sering digunakan di golang

// type data slice adalah potongan dari data array
// slice mirip dengan Array, yg membedakan adalah ukuran slice bisa berubah
// slice dan array selalu terkoneksi, dimana slice adalah data yg mengakses sebagian atau seluruh data di array

// kalau data array sudah di tentukan misalnya data array nya adalah 5 ya, itu jadi 5 berbeda dengan slice yang data nya bisa bertambah secara dinamis

// * Detail tipe data slice
// - tipe data slice memiliki 3 data, yaitu pointer, length, dan capacity
// - pointer adalah penunjuk data pertama di array pada slice
// - length adalah panjang dari slice,
// - capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity

// * Membuat Slice dari Array

// Membuat Slice | Keterangan
// array[low:hight] | Membuat Slice dari Array dimulai index low sampai index sebelum high
// array[low:] | Membuat Slice dari Array dimulai dari index low sampai index akhir di Array
// array[:high] | Membuat Slice dari Array dimulai dari index 0 sampai index sebelum high
// array[:] | Membuat Slice dari Array dimulai dari index 0 sampai index akhir di Array

// Slice dan Array
//
// contoh :
//
// array{
//   0 = "January",
//   1 = "February",
//   2 = "Maret",
//   3 = "April",
//   4 = "Mei",
//   5 = "Juni",
//   6 = "Juli",
//   7 = "Agustus",
//   8 = "September",
//   9 = "Oktober",
//   10 = "November",
//   11 = "Desember",
// }
//
// array[4:7]
//
// - slice 1 :
// -- pointer = 4 => index ke 4 => penunjuk data pertamanya
// -- length = 3 => panjang data dimulai dari data pertama, karena sebelum high maka dari data 4 5 6 = 3 total length nya jadi dari 4 - 7 itu jarrak nya 3
// -- capacity = 8 => total semua kapasitas dari array nya dimulai dari index ke 4 data pertama(pointer)
//
// contoh lainnya
//
// array[6:9]
//
// - slice 2 :
// -- pointer = 6
// -- length = 3 => karena dari 6 ke 9 itu jaraknya 3
// -- capacity = 6 => karena dari data pertama yaitu index ke 6 hingga ke index data terakhir, karena array nya berisi sampai index ke 11 maka dari 6 ke 11 itu adalah 6 7 8 9 10 11
//

// # Kode Program Slice

package main

import "fmt"

func main() {
	names := [...]string{"Rahman", "Visalux", "Gaming", "Riko", "Risma", "Jancok"}
	slice := names[4:6]

	fmt.Println(slice[0])
	fmt.Println(slice[1])
}
