package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"power4/grid"
	"power4/structure"
	"power4/utils"
	"strconv"
)

// renderTemplate est une fonction utilitaire pour afficher un template HTML avec des données dynamiques
func renderTemplate(w http.ResponseWriter, filename string, data interface{}) {
	tmpl := template.Must(template.ParseFiles("template/" + filename))
	tmpl.Execute(w, data)
}

// Home gère la page d'accueil
func Home(w http.ResponseWriter, r *http.Request) {

	*grid.IsRetrievePointer = false

	data := structure.One{
		Title:    "Puissance 4",
		Message:  "Bienvenue sur le jeu du Puissance 4 ! Vous pouvez commencer une nouvelle partie ou continuer une partie sauvegardée. Amusez-vous bien !",
		Historic: []structure.Partie{{Date: "1", Joueur1: "martin", Joueur2: "kevin", Winner: "martin"}, {Date: "2", Joueur1: "martin", Joueur2: "kevin", Winner: "martin"}, {Date: "3", Joueur1: "kevin", Joueur2: "martin", Winner: "kevin"}, {Date: "1", Joueur1: "martin", Joueur2: "kevin", Winner: "martin"}, {Date: "2", Joueur1: "martin", Joueur2: "kevin", Winner: "martin"}, {Date: "1", Joueur1: "martin", Joueur2: "kevin", Winner: "martin"}, {Date: "2", Joueur1: "martin", Joueur2: "kevin", Winner: "martin"}, {Date: "1", Joueur1: "martin", Joueur2: "kevin", Winner: "martin"}, {Date: "2", Joueur1: "martin", Joueur2: "kevin", Winner: "martin"}},
	}

	renderTemplate(w, "home.html", data) // Affiche le template index.html avec les données
}

func Save(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	homeData := structure.One{
		Title:    "Partie enregistré !",
		Message:  "Bienvenue sur le jeu du Puissance 4 ! Vous pouvez commencer une nouvelle partie ou continuer une partie sauvegardée. Amusez-vous bien !",
		Message2: "Ici, vous retrouvez les ancienne partie jouer :",
		Historic: []structure.Partie{{Date: "1", Joueur1: "martin", Joueur2: "kevin"}, {Date: "2", Joueur1: "martin", Joueur2: "kevin"}, {Date: "3", Joueur1: "kevin", Joueur2: "martin"}},
	}

	game := structure.GameData{
		Player1: 1,
		Player2: 2,
		Grid:    *grid.PointerGrid,
		Turn:    1,
		IsOver:  false,
	}

	*grid.IsRetrievePointer = false

	saveJSON("gamesave.json", game)

	renderTemplate(w, "home.html", homeData)

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

// gestion de la grille
func Game(w http.ResponseWriter, r *http.Request) {
	title := "Pion à poser : "
	playerTurn := *(grid.PlayerTurnPointer)
	visibility := "auto"
	winner := "dimitri"
	textvisibility := "none"

	if r.Method == http.MethodPost {
		col := r.FormValue("col")
		colInt, _ := strconv.Atoi(col)
		player, iswon := grid.SetToken(colInt)

		if iswon && player != 0 {
			//pointers = "none"
			visibility = "none"
			textvisibility = "auto"
			winner = "Felications joueur " + strconv.Itoa(player)
			addGameToHistoric(playerTurn, 0, player, "14/10/2025")
		}

		if iswon && player == 0 {
			visibility = "none"
			textvisibility = "auto"
			winner = "Egalité"
			addGameToHistoric(playerTurn, 0, 0, "14/10/2025")
		}
	}

	data := structure.PageData{
		Title:          title,
		Grid:           *grid.PointerGrid,
		PlayerTurn:     playerTurn,
		Color:          grid.SetColor(),
		Visibility:     visibility,
		Winner:         winner,
		TextVisibility: textvisibility,
	}

	renderTemplate(w, "play.html", data)
}

func Returnmenu(w http.ResponseWriter, r *http.Request) {
	utils.EmptyGrid()
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func Reset(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	utils.EmptyGrid()

	http.Redirect(w, r, "/play", http.StatusSeeOther)
}

func Replay(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	utils.EmptyGrid()
	*grid.PlayerTurnPointer = 1

	http.Redirect(w, r, "/play", http.StatusSeeOther)
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

func saveJSON(nomFichier string, data interface{}) error {
	// Convertir les données en JSON (avec indentation pour la lisibilité)
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// Écrire dans le fichier
	return os.WriteFile(nomFichier, bytes, 0644)
}

func addGameToHistoric(player1 int, player2 int, winner int, date string) {
	var historic []structure.Historic

	file, err := os.ReadFile("gamehistoric.json")
	if err == nil {
		if err := json.Unmarshal(file, &historic); err != nil {
			fmt.Println("Erreur lors du décodage JSON :", err)
		}
	} else if !os.IsNotExist(err) {
		fmt.Println("Erreur de lecture fichier :", err)
	}

	newGame := structure.Historic{
		Player1: player1,
		Player2: player2,
		Winner:  winner,
		Date:    date,
	}
	historic = append(historic, newGame)

	saveJSON("gamehistoric.json", historic)
}
