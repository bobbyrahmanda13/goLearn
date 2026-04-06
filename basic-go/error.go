package main

import (
	"errors"
	"fmt"
)

var errorBro = errors.New("Pembagian dengan 0")

func Pembagian(nilai int, pembagian int) (int, error) {
	if pembagian == 0 {
		return 0, errorBro
	} else {
		return nilai / pembagian, nil
	}
}

func main() {
	hasil, err := Pembagian(20, 0)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
