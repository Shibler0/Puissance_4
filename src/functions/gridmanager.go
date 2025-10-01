package functions

import "fmt"

func Gridmanager(column int, grid *[6][7]int, player *string) {

	row, isEmpty := getEmptyCase(grid, column)

	if isEmpty {
		return
	}
	fmt.Printf("joueur %s a choisi %d \n", string(*player), column)
	fmt.Printf("Pion placé a la %d eme ligne \n", row)
}
