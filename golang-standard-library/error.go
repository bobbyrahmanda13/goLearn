package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not fount error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "bobby" {
		return NotFoundError
	}

	// sukses
	return nil
}

func main() {
	err := GetById("windy")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("not found error")
		} else {
			fmt.Println("unknown error")
		}
	}
}
