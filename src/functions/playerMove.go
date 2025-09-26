package functions

import "fmt"

func PlayerMoove(grid *[6][7]int, player *string) {

	var column int

	fmt.Println("Choisissez votre column de 1 a 7")
	fmt.Scan(&column)

	row, isFull := getEmptyCase(grid, column)

	for column > 7 || column < 0 || isFull {
		fmt.Println("Choisissez votre column uniquement de 1 a 7")
		fmt.Scan(&column)
		row, isFull = getEmptyCase(grid, column)
	}

	grid[row][column] = 1 //place le pion

	fmt.Printf("joueur %s a choisi %d \n", string(*player), column)
	fmt.Printf("Pion placé dans %d \n", row)

}

func getEmptyCase(grid *[6][7]int, column int) (int, bool) {

	for row := len(grid) - 1; row >= 0; row-- { //verifie si la case est vide en partant du bas
		if grid[row][column] == 0 { // si la case est vide alors renvoie celle-ci
			return row, false
		}
	}

	return -1, true
}

func gridScan(grid *[6][7]int) {

}
