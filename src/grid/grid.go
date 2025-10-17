package grid

import "math/rand"

var grid [6][7]int
var PointerGrid = &grid

var playerTurn int = 1 // 1 = jaune, 2 = rouge
var PlayerTurnPointer = &playerTurn

var numberOfPlays = 0
var numberOfPlaysPointer = &numberOfPlays

var IsRetrieve bool = true
var IsRetrievePointer = &IsRetrieve

// création de la grille
func CreateGrid() {
	for i := range 6 {
		for j := range 7 {
			PointerGrid[i][j] = 0 // Remplissage du tableau
		}
	}
}

// mettre les jetons
func SetToken(column int) (int, bool) {

	if *IsRetrievePointer {

		for row := len(grid) - 1; row >= 0; row-- { //verifie si la case est vide en partant du bas
			if PointerGrid[row][column] == 0 { // si la case est vide alors renvoie celle-ci
				PointerGrid[row][column] = *PlayerTurnPointer
				*numberOfPlaysPointer++
				player, iswin := checkWinCondition(row, column)
				managePlayerTurn()
				return player, iswin
			}
		}

		return 0, false
	} else {
		*IsRetrievePointer = true
		return 0, false
	}
}

func RandomEncouragement() string {

	tab := []string{
		"Tu vas le faire", "c'est pas fini", "Bon courage",
	}

	n := rand.Intn(3)

	return tab[n]
}

// choisi quel joueur doit jouer
func managePlayerTurn() {
	if *PlayerTurnPointer == 1 {
		*PlayerTurnPointer = 2
		return
	}
	if *PlayerTurnPointer == 2 {
		*PlayerTurnPointer = 1
		return
	}
}

func SetColor() string {
	if *PlayerTurnPointer == 2 {
		return "red"
	} else {
		return "yellow"
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

	//verifie si une personen est gagnante
	for i := 1; i < 4; i++ {

		if column+i < 7 && PointerGrid[row][column+i] == *PlayerTurnPointer && isRowAligned1 {
			row1++
		} else {
			isRowAligned1 = false
		}
		if column+i < 7 && row+i < 6 && PointerGrid[row+i][column+i] == *PlayerTurnPointer && isRowAligned2 {
			diagonalDroite++
		} else {
			isRowAligned2 = false
		}
		if row+i < 6 && PointerGrid[row+i][column] == *PlayerTurnPointer && isRowAligned3 {
			vertical++
		} else {
			isRowAligned3 = false
		}
		if row+i < 6 && column-i >= 0 && PointerGrid[row+i][column-i] == *PlayerTurnPointer && isRowAligned4 {
			diagonalGauche++
		} else {
			isRowAligned4 = false
		}
		if column-i >= 0 && PointerGrid[row][column-i] == *PlayerTurnPointer && isRowAligned5 {
			row1++
		} else {
			isRowAligned5 = false
		}
		if column-i >= 0 && row-i >= 0 && PointerGrid[row-i][column-i] == *PlayerTurnPointer && isRowAligned6 {
			diagonalDroite++
		} else {
			isRowAligned6 = false
		}
		if row-i >= 0 && PointerGrid[row-i][column] == *PlayerTurnPointer && isRowAligned7 {
			vertical++
		} else {
			isRowAligned7 = false
		}
		if row-i >= 0 && column+i < 7 && PointerGrid[row-i][column+i] == *PlayerTurnPointer && isRowAligned8 {
			diagonalGauche++
		} else {
			isRowAligned8 = false
		}
	}

	var victory = false

	if row1 >= 3 || vertical >= 3 || diagonalDroite >= 3 || diagonalGauche >= 3 {
		victory = true
		return *PlayerTurnPointer, true
	}

	var isGridFull = true

	for row := 0; row < 6; row++ {
		for col := 0; col < 7; col++ {
			if PointerGrid[row][col] == 0 {
				isGridFull = false
				break
			}
		}
		if !isGridFull {
			break
		}
	}

	if isGridFull && !victory {
		return 0, true
	}

	return 0, false

}
