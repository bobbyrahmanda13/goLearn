package main

import (
	"fmt"
	"time"
)

/* var (
	nama string
	umur int
)

fmt.Print("Masukkan Nama : ")
fmt.Scanln(&nama)
fmt.Print("Masukkan Umur : ")
fmt.Scanln(&umur)

fmt.Println("Halo", nama, "Umur kamu", umur, "tahun") */

func typeWrite(text string, delay time.Duration) {
	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(delay)
	}
	fmt.Println()
}

func main() {
	typeWrite("Hai Sayangg", 70*time.Millisecond)

	fmt.Println("")
	typeWrite("Setelah kamu sama abg apakah kamu ada bahagia atau tidak ?", 70*time.Millisecond)
	fmt.Println("")
	typeWrite("Jika kamu bahagia hanya Alhamdulillah yang bisa abg jawab, jika tidak abg minta maaf atas sifat dan kekurangan abg tapi abg berharap kamu menerima serta bantu abg memperbaikinya", 70*time.Millisecond)
	fmt.Println("")
	typeWrite("Sayangg", 70*time.Millisecond)
	fmt.Println("")
	typeWrite("Abg sangat Mencintaimu, apakah kita bisa lalui jalan ini Sayang ?", 70*time.Millisecond)
	fmt.Println("")
	typeWrite("Apakah kita bisa langgeng sampai menikah nanti", 70*time.Millisecond)
	fmt.Println("")
	typeWrite("Abg sangat ingin menikahimu, harus seperti apakah abg berjuang untuk mu Sayang", 70*time.Millisecond)
	fmt.Println("")
	typeWrite("Abg ga tau ntah kenapa sewaktu kita berantem kadang perasaan atas kamu hilang sebentar, tapi hanya sebentar trus jatuh cinta lagi sama kamu jika lihat kamu terseyum", 70*time.Millisecond)
	fmt.Println("")
	typeWrite("Abg Sayang Kamu, Abg Kangen Kamu, Abg Rindu Kamu, Serasa pengen peluk saja, untuk melepas rasa Rindu ini Sayang", 70*time.Millisecond)
	fmt.Println("")
	typeWrite("I LOVE YOU SAYANG, SARANGHAEE ", 70*time.Millisecond)
	fmt.Println("")
}
