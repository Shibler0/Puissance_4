package main

import (
	"fmt"
	"net/http"
	"power4/grid"
	"power4/router"
)

func main() {
	// Charge le routeur
	r := router.New()

	//crée une grille de jeu vide
	grid.CreateGrid()

	//choisi la route de départ
	fmt.Println("🚀 Serveur démarré sur http://localhost:8080/play")
	//crée le serveur
	http.ListenAndServe(":8080", r)

}
