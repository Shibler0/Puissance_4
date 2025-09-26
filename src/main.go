package main

import (
	"fmt"
	"power4/functions"
)

func main() {
	// Charge le routeur
	// r := router.New()

	// fmt.Println("🚀 Serveur démarré sur http://localhost:8080")
	// http.ListenAndServe(":8080", r)

	fmt.Println("Puissance 4")
	player1 := functions.EnterPseudo()
	player2 := functions.EnterPseudo()

	grid := [6][7]int{}
	ptrGrid := &grid

	for !functions.IsGameover() {

		var currentPlayer *string = functions.PlayerTurn(&player1, &player2) //jongle entre 2 les 2 joueurs
		functions.PlayerMoove(ptrGrid, currentPlayer)                        //demande au joueur d'entrer son numero de column
	}

}
