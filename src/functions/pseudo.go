package functions

import "fmt"

func EnterPseudo() string {

	var name string

	fmt.Println("Entrez votre nom")

	fmt.Scan(&name)

	return name
}
