package router

import (
	"net/http"
	"power4/controller"
)

// New crée et retourne un nouvel objet ServeMux configuré avec les routes de l'application
func New() *http.ServeMux {

	mux := http.NewServeMux() // Création d'un nouveau ServeMux, qui est un routeur simple pour les requêtes HTTP

	// On associe les chemins URL à des fonctions spécifiques du controller
	mux.HandleFunc("/home", controller.Home)
	mux.HandleFunc("/play", controller.Game)
	mux.HandleFunc("/save", controller.Save)
	mux.HandleFunc("/returnmenu", controller.Returnmenu)
	mux.HandleFunc("/reset", controller.Reset)
	mux.HandleFunc("/replay", controller.Replay)

	//gere le css
	fileServer := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	return mux // On retourne le routeur configuré
}
