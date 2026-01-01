package main

import "fmt"

func fizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		result := ""

		if i%3 == 0 {
			result += "Fizz"
		}
		if i%5 == 0 {
			result += "Buzz"
		}

		if result == "" {
			fmt.Println(i)
		} else {
			fmt.Println(result)
		}
	}
}

func main() {
	fizzBuzz(100)
}
