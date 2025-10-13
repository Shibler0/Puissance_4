package router

import (
	"net/http"
	"power4/controller"
)

// New crée et retourne un nouvel objet ServeMux configuré avec les routes de l'application
func New() *http.ServeMux {

	mux := http.NewServeMux() // Création d'un nouveau ServeMux, qui est un routeur simple pour les requêtes HTTP

	// On associe les chemins URL à des fonctions spécifiques du controller
	mux.HandleFunc("/home", controller.Home) // "/" correspond à la page d'accueil. Appelle la fonction Home du controller
	mux.HandleFunc("/play", controller.Game)
	//mux.HandleFunc("/play", controller.HandleGrid)
	//mux.HandleFunc("/contact", controller.Contact) // "/contact" correspond à la page de contact. Appelle la fonction Contact

	fileServer := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	return mux // On retourne le routeur configuré
}
