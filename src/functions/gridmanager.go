package functions

import "fmt"

func Gridmanager(column int, grid *[6][7]int, player *string) {

	row, isEmpty := getEmptyCase(grid, column)

	if isEmpty {
		return
	}
	fmt.Printf("joueur %s a choisi %d \n", string(*player), column)
	fmt.Printf("Pion placé dans %d", row)
}

func getEmptyCase(grid *[6][7]int, column int) (int, bool) {

	for row := len(grid) - 1; row >= 0; row-- { //verifie si la case est vide en partant du bas
		if grid[row][column] == 0 { // si la case est vide alors renvoie celle-ci
			return row, false
		}
	}

	return -1, true
}

// func isFull(grid *[6][7]int, column int) bool {

// 	for row := len(grid) - 1; row >= 0; row-- { //verifie si la column est remplie
// 		if grid[row][column] == 0 {
// 			return false
// 		}
// 	}

// 	return true

// }
