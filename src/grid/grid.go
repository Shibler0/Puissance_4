package grid

import (
	"fmt"
)

type PageData struct {
	Title string
	Grid  [6][7]int
}

var grid [6][7]int
var PointerGrid = &grid

var playerTurn int = 1 // 1 = jaune, 2 = rouge
var playerTurnPointer = &playerTurn

var numberOfPlays = 0
var numberOfPlaysPointer = &numberOfPlays

// création de la grille
func CreateGrid() {

	for i := range 6 {
		for j := range 7 {
			PointerGrid[i][j] = 0 // Valeur visible
		}
	}

	//Debug pour check si la grille est juste
	fmt.Println("Grille envoyée au template :", PointerGrid)
}

// mettre les jetons
func SetToken(row int, column int) {
	for row := len(grid) - 1; row >= 0; row-- { //verifie si la case est vide en partant du bas
		if PointerGrid[row][column] == 0 { // si la case est vide alors renvoie celle-ci
			PointerGrid[row][column] = *playerTurnPointer
			*numberOfPlaysPointer++
			checkWinCondition(row, column)
			managePlayerTurn()
			return
		}
	}
}

// choisi quel joueur doit jouer
func managePlayerTurn() {
	if *playerTurnPointer == 1 {
		*playerTurnPointer = 2
		return
	}
	if *playerTurnPointer == 2 {
		*playerTurnPointer = 1
		return
	}
}

func checkWinCondition(row int, column int) (int, bool) {

	var row1 int = 0
	var diagonalDroite int = 0
	var vertical int = 0
	var diagonalGauche int = 0

	isRowAligned1 := true
	isRowAligned2 := true
	isRowAligned3 := true
	isRowAligned4 := true
	isRowAligned5 := true
	isRowAligned6 := true
	isRowAligned7 := true
	isRowAligned8 := true

	if *numberOfPlaysPointer < 7 {
		return 0, false
	}

	for i := 0; i < 7; i++ {
		if PointerGrid[0][i] != 0 {
		}
	}

	for i := 1; i < 4; i++ {

		if column+i < 7 && PointerGrid[row][column+i] == *playerTurnPointer && isRowAligned1 {
			row1++
		} else {
			isRowAligned1 = false
		}
		if column+i < 7 && row+i < 6 && PointerGrid[row+i][column+i] == *playerTurnPointer && isRowAligned2 {
			diagonalDroite++
		} else {
			isRowAligned2 = false
		}
		if row+i < 6 && PointerGrid[row+i][column] == *playerTurnPointer && isRowAligned3 {
			vertical++
		} else {
			isRowAligned3 = false
		}
		if row+i < 6 && column-i >= 0 && PointerGrid[row+i][column-i] == *playerTurnPointer && isRowAligned4 {
			diagonalGauche++
		} else {
			isRowAligned4 = false
		}
		if column-i >= 0 && PointerGrid[row][column-i] == *playerTurnPointer && isRowAligned5 {
			row1++
		} else {
			isRowAligned5 = false
		}
		if column-i >= 0 && row-i >= 0 && PointerGrid[row-i][column-i] == *playerTurnPointer && isRowAligned6 {
			diagonalDroite++
		} else {
			isRowAligned6 = false
		}
		if row-i >= 0 && PointerGrid[row-i][column] == *playerTurnPointer && isRowAligned7 {
			vertical++
		} else {
			isRowAligned7 = false
		}
		if row-i >= 0 && column+i < 7 && PointerGrid[row-i][column+i] == *playerTurnPointer && isRowAligned8 {
			diagonalGauche++
		} else {
			isRowAligned8 = false
		}
	}

	if row1 >= 3 || vertical >= 3 || diagonalDroite >= 3 || diagonalGauche >= 3 {
		fmt.Println("Gagné !")
		return *playerTurnPointer, true
	}

	return 0, false

}
