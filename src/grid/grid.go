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
