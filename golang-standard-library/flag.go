// -- Package Flag
// -- Package flag berisikan fungsionalitas untuk memparsing command line argument
package main

import (
	"flag"
	"fmt"
)

func main() {
	var username *string = flag.String("username", "root", "database username")
	var password *string = flag.String("password", "root", "database password")
	var host *string = flag.String("host", "localhost", "database host")
	var port *int = flag.Int("port", 0, "database port")

	flag.Parse() // untuk melakukan parsing

	fmt.Println("Username ", *username)
	fmt.Println("Password ", *password)
	fmt.Println("host ", *host)
	fmt.Println("port ", *port)
}
