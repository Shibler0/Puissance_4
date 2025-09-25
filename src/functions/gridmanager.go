package functions

import "fmt"

func Gridmanager(column int, grid *[6][7]int, player *string) {
	if isFill() {
		return
	}
	fmt.Printf("joueur %s a choisi %d \n", string(*player), column)
}

func isFill() bool {
	return false
}
