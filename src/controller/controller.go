package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"power4/grid"
	"power4/structure"
	"strconv"
)

// renderTemplate est une fonction utilitaire pour afficher un template HTML avec des données dynamiques
func renderTemplate(w http.ResponseWriter, filename string, data interface{}) {
	tmpl := template.Must(template.ParseFiles("template/" + filename))
	tmpl.Execute(w, data)
}

// Home gère la page d'accueil
func Home(w http.ResponseWriter, r *http.Request) {
	data := structure.One{
		Title:   "Puissance 4",
		Message: "Bienvenue sur le jeu",
	}
	renderTemplate(w, "home.html", data) // Affiche le template index.html avec les données
}

func Save(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	game := structure.GameData{
		Player1: 1,
		Player2: 2,
		Grid:    *grid.PointerGrid,
		Turn:    1,
		IsOver:  false,
	}

	enregistrerJSON("save.json", game)

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

// gestion de la grille
func Game(w http.ResponseWriter, r *http.Request) {
	title := "Pion à poser : "
	playerTurn := *(grid.PlayerTurnPointer)
	visibility := "auto"
	winner := "dimitri"

	if r.Method == http.MethodPost {
		col := r.FormValue("col")
		colInt, _ := strconv.Atoi(col)
		player, iswon := grid.SetToken(colInt)
		if iswon {
			//pointers = "none"
			visibility = "none"
			winner = strconv.Itoa(player)
			fmt.Printf("%d a gagné", player)
		}
	}

	data := structure.PageData{
		Title:      title,
		Grid:       *grid.PointerGrid,
		PlayerTurn: playerTurn,
		Color:      grid.SetColor(),
		Visibility: visibility,
		Winner:     winner,
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

func enregistrerJSON(nomFichier string, data interface{}) error {
	// Convertir les données en JSON (avec indentation pour la lisibilité)
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// Écrire dans le fichier
	return os.WriteFile(nomFichier, bytes, 0644)
}
