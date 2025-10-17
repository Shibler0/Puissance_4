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

	//*grid.IsRetrievePointer = false

	utils.LoadJSON()

	renderTemplate(w, "home.html", utils.LoadJSON()) // Affiche le template index.html avec les données
}

func Save(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	game := structure.GameData{
		Player1: *structure.PointerPlayer1,
		Player2: *structure.PointerPlayer2,
		Grid:    *grid.PointerGrid,
		Turn:    1,
		IsOver:  false,
	}

	//*grid.IsRetrievePointer = false

	saveJSON("gamesave.json", game)

	renderTemplate(w, "home.html", utils.LoadJSON())

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

// gestion de la grille
func Game(w http.ResponseWriter, r *http.Request) {
	title := "Pion à poser : "
	playerTurn := *(grid.PlayerTurnPointer)
	visibility := "auto"
	winner := "dimitri"
	textvisibility := "none"
	encouragement := "Bonne chance"

	var x *string = structure.PointerPlayer1
	var y *string = structure.PointerPlayer2

	if *x == "" && *y == "" {
		*x = "Joueur1"
		*y = "Joueur2"
	}

	if *(grid.PlayerTurnPointer) == "" {
		*(grid.PlayerTurnPointer) = *x
		*(grid.ColorPointer) = 1
	}

	if r.Method == http.MethodPost {
		formPlayer1 := r.FormValue("Joueur1")
		formPlayer2 := r.FormValue("Joueur2")

		if formPlayer1 != "" {
			*structure.PointerPlayer1 = formPlayer1
		}
		if formPlayer2 != "" {
			*structure.PointerPlayer2 = formPlayer2
		}

		encouragement = grid.RandomEncouragement()
		col := r.FormValue("col")

		if col != "" {
			colInt, _ := strconv.Atoi(col)
			player, iswon := grid.SetToken(colInt)

			if iswon && player != 0 {
				visibility = "none"
				textvisibility = "auto"
				winner = "Félicitations joueur " + strconv.Itoa(player)
				addGameToHistoric(*x, *y, player, "16/10/2025")
			}

			if iswon && player == 0 {
				visibility = "none"
				textvisibility = "auto"
				winner = "Égalité"
				addGameToHistoric(*x, *y, 0, "16/10/2025")
			}
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
		Encouragement:  encouragement,
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
	*grid.ColorPointer = 1

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
	// Convertir les données en JSON
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// Écrire dans le fichier
	return os.WriteFile(nomFichier, bytes, 0644)
}

func addGameToHistoric(player1 string, player2 string, winner int, date string) {
	var historic []structure.Historic

	fmt.Println("3", player1, player2, winner, date)

	file, err := os.ReadFile("gamehistoric.json")
	if err == nil {
		if err := json.Unmarshal(file, &historic); err != nil {
			fmt.Println("Erreur lors du décodage JSON :", err)
		}
	} else if !os.IsNotExist(err) {
		fmt.Println("Erreur de lecture fichier :", err)
	}

	fmt.Println("4", &player1, player2, winner, date)

	newGame := structure.Historic{
		Player1: player1,
		Player2: player2,
		Winner:  winner,
		Date:    date,
	}
	fmt.Println("5", player1, player2, winner, date)
	fmt.Println(newGame.Player1, newGame.Player2, newGame.Winner, newGame.Date)
	historic = append(historic, newGame)

	saveJSON("gamehistoric.json", historic)

	fmt.Println("6", player1, player2, winner, date)
}
