/* # Interface Kosong
- Go-lang bukanlah bahasa pemrograman yang berioentasi Objek
- biasanya dalam pemrograman berioentasi Objek, ada satu data parent di puncak yang bisa dianggap sebagai semua implementasi data yang ada di bahasa pemrograman tersebut
- contoh di Java ada java.lang.Object
- untuk menangani kasus seperti ini, di Go-lang kita bisa menggunakan interface kosong
- interface kosong adalah interface yang tidak memiliki deklarasi method satupun, hal ini membuat secara otomatis semua tipe data akan menjadi implementasinya
- interface kosong, juga memiliki type alias bernama any */

/* Penggunaan Interface Kosong di Go-Lang

- ada banyak contoh penggunaan interface kosong di Go-Lang, seperti :
 - fmt.Println(a ...interface{})
 - panic(v interface{})
 - recover() interface{}
 - dan lain2 */

package main

import "fmt"

func Ups() any {
	// return 1
	// return true
	return "Ups"
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}
