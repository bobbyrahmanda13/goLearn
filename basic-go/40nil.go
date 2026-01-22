/* # Nil
- biasanya didalam bahasa pemrograman lain, object yang belum diinisialisasi maka secara otomatis nilainya adalah null atau nill
- berbeda dengan Go-Lang, di Go-Lang saat kita buat variable dengan tipe data tertentu, maka secara otomatis akan dibuatkan default value nya
- Namun di Go-Lang ada data nil, yaitu data kosong
- Nil sendiri hanya bisa digunakan dibeberapa tipe data, seperti interface, function, map, slice, pointer dan channel */

package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	// data := NewMap("Bobby")
	data := NewMap("")
	if data == nil {
		fmt.Println("Data map masih kosong")
	} else {
		fmt.Println(data["name"])
	}
}
