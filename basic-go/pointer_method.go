package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	bobby := Man{"bobby"}
	bobby.Married()

	fmt.Println(bobby.Name)
}
