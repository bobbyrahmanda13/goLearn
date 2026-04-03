package helper

import "fmt"

var (
	version     = "1.0.0"
	Application = "golang"
)

func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}

func Contoh() {
	result := sayGoodBye("Bobby")
	fmt.Println(result)
	fmt.Println(version)
}
