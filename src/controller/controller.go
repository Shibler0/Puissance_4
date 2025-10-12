package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"power4/grid"
	"strconv"
)

// renderTemplate est une fonction utilitaire pour afficher un template HTML avec des données dynamiques
func renderTemplate(w http.ResponseWriter, filename string, data interface{}) {
	tmpl := template.Must(template.ParseFiles("template/" + filename))
	tmpl.Execute(w, data)
}

// Home gère la page d'accueil
func Home(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title":   "Accueil",                           // Titre de la page
		"Message": "Bienvenue sur la page d'accueil 🎉", // Message affiché dans le template
	}
	renderTemplate(w, "index.html", data) // Affiche le template index.html avec les données
}

// gestion de la grille
func Game(w http.ResponseWriter, r *http.Request) {
	title := "Grille Puissance 4"

	if r.Method == http.MethodPost {
		col := r.FormValue("col")
		colInt, _ := strconv.Atoi(col)
		title = fmt.Sprintf("Colonne sélectionnée : %d", colInt)
		grid.SetToken(5, colInt)
	}

	data := grid.PageData{
		Title: title,
		Grid:  *grid.PointerGrid,
	}

	renderTemplate(w, "play.html", data)
}

// Contact gère la page de contact
func Contact(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost { // Si le formulaire est soumis en POST
		// Récupération des données du formulaire
		name := r.FormValue("name") // Récupère le champ "name"
		msg := r.FormValue("msg")   // Récupère le champ "msg"

		data := map[string]string{
			"Title":   "Contact",
			"Message": "Merci " + name + " pour ton message : " + msg, // Message personnalisé après soumission
		}
		renderTemplate(w, "contact.html", data)
		return // On termine ici pour ne pas exécuter la partie GET
	}

	// Si ce n'est pas un POST, on affiche simplement le formulaire
	data := map[string]string{
		"Title":   "Contact",
		"Message": "Envoie-nous un message 📩",
	}
	renderTemplate(w, "contact.html", data)
}
