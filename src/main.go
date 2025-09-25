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

	isWon := false

	for !isWon {

		functions.IsGameover()

		var pseudo = "dimitri"

		fmt.Println("Puissance 4")
		fmt.Println("Entrez votre nom")
		fmt.Scan()

		functions.Gridmanager(0, pseudo)
	}
}
