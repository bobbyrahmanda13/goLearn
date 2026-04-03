package main

import (
	"fmt"

	"basic-go/database"
	_ "basic-go/internal"
)

func main() {
	fmt.Println(database.GetDatabase())
}
