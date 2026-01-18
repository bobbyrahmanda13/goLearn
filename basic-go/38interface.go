/*
# Interface
- adalah tipe data abstract, dia tidak memiliki implementasi lansung
- sebuah interface berisikan definisi-definisi method
- biasanya interface digunakan sebagai kontrak
*/

package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Eko"}
	SayHello(person)

	animal := Animal{Name: "Puti Kucing The Best"}
	SayHello(animal)
}
