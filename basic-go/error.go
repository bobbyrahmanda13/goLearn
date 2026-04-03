package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagian int) (int, error) {
	if pembagian == 0 {
		return 0, errors.New("Pembagian dengan 0")
	} else {
		return nilai / pembagian, nil
	}
}

func main() {
	hasil, err := Pembagian(20, 5)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
