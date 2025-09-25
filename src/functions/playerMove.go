package functions

import "fmt"

func PlayerMoove(grid *[6][7]int, player *string) int {

	var choice int

	fmt.Println("Choisissez votre column de 1 a 7")
	fmt.Scan(&choice)

	return choice

}
